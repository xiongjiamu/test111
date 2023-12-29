package dkdnserver

import (
	"dkdns/dkFramework/configs"
	"dkdns/dkFramework/logger"
	"fmt"
	"github.com/miekg/dns"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type CachedResponse struct {
	Msg      *dns.Msg
	ExpireAt time.Time
}

var (
	cache      map[string]CachedResponse
	cacheMutex sync.RWMutex
)

func lookupFromCache(queryKey string) (*dns.Msg, bool) {
	cacheMutex.RLock()
	defer cacheMutex.RUnlock()
	cachedResponse, found := cache[queryKey]
	if !found || time.Now().After(cachedResponse.ExpireAt) {
		return nil, false
	}
	return cachedResponse.Msg.Copy(), true
}

func storeInCache(queryKey string, response *dns.Msg, timeout time.Duration) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	cache[queryKey] = CachedResponse{
		Msg:      response.Copy(),
		ExpireAt: time.Now().Add(timeout),
	}
}

func isIPInRange(ip string, ipRanges []*net.IPNet) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	for _, ipNet := range ipRanges {
		if ipNet.Contains(parsedIP) {
			return true
		}
	}

	return false
}

func loadIPRangesFromFile(filepath string) ([]*net.IPNet, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var ipRanges []*net.IPNet
	ips := strings.Split(string(data), "\n")
	for _, ip := range ips {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			_, ipNet, err := net.ParseCIDR(ip)
			if err != nil {
				fmt.Println("Error parsing IP range:", err)
				continue
			}
			ipRanges = append(ipRanges, ipNet)
		}
	}

	return ipRanges, nil
}
func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg, appConfigs *configs.Config) {
	m := new(dns.Msg)
	m.SetReply(r)

	ipRanges, err := loadIPRangesFromFile(appConfigs.DNS.Whitelist)
	if err != nil {
		logger.Println("Error loading IP ranges:", err)
		return
	}

	for _, q := range r.Question {
		name := strings.ToLower(q.Name)
		if q.Qtype == dns.TypeA {
			queryKey := fmt.Sprintf("%s-%d", name, q.Qtype)
			cachedResponse, found := lookupFromCache(queryKey)
			if found {
				logger.Println("Cache hit for:", name)

				resp := new(dns.Msg)
				resp.SetReply(r)
				resp.Answer = cachedResponse.Answer // Copy answers from cached response

				resp.Id = r.Id // Set the response ID to match the query ID

				w.WriteMsg(resp)
				return
			}

			ip, err := net.LookupIP(name)
			if err == nil {
				for _, addr := range ip {
					ipStr := addr.String()
					if isIPInRange(ipStr, ipRanges) {
						rr := &dns.A{
							Hdr: dns.RR_Header{Name: name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
							A:   addr,
						}
						m.Answer = append(m.Answer, rr)
						logger.Println("localDNS:", m.String())
						storeInCache(queryKey, m, appConfigs.DNS.CacheTimeout) // Cache the response
						w.WriteMsg(m)
						return
					}
				}
			}

			recursiveAddr := appConfigs.DNS.RecursiveDNS
			c := new(dns.Client)
			resp, _, err := c.Exchange(r, recursiveAddr+":53")
			if err != nil {
				logger.Println("Error:", err)
				return
			}
			logger.Println("recursiveDNS:", resp.String())
			storeInCache(queryKey, resp, appConfigs.DNS.CacheTimeout) // Cache the response
			w.WriteMsg(resp)
		}
	}
}

func RunDNSServer(appConfigs *configs.Config) {

	udpHandler := dns.NewServeMux()
	udpHandler.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		handleDNSRequest(w, r, appConfigs)
	})

	cache = make(map[string]CachedResponse) // Initialize the cache
	server := &dns.Server{Addr: ":53", Net: "udp", Handler: udpHandler}

	logger.Println("Starting DNS server at port 53...")
	if err := server.ListenAndServe(); err != nil {
		logger.Println("Error starting DNS server:", err)
		os.Exit(1)
	}

	// 优雅地关闭服务器
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Println("Shutting down DNS server")
	server.Shutdown()
}
