package dkdnserver

import (
	"dkdns/dkFramework/configs"
	"dkdns/dkFramework/logger"
	special_list "dkdns/httpServer/services"
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

func GetCacheMap() (cacheMap map[string]CachedResponse) {
	return cache
}

func isPrivateIP(ip net.IP) bool {
	cidrs := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	for _, cidr := range cidrs {
		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(err)
		}
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}

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

// 删除指定 key 的缓存项
func DeleteFromCache(queryKey string) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	delete(cache, queryKey)
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
	startTime := time.Now() // 记录开始处理请求的时间
	// 获取客户端IP
	clientIP, _, _ := net.SplitHostPort(w.RemoteAddr().String())

	ipRanges, err := loadIPRangesFromFile(appConfigs.DNS.Whitelist)
	if err != nil {
		logger.Println("Error loading IP ranges:", err)
		return
	}

	for _, q := range r.Question {
		name := strings.ToLower(q.Name)
		if q.Qtype == dns.TypeA {
			queryKey := fmt.Sprintf("%s-%d", name, q.Qtype)
			/**
			先查询 Special_list ,
			*/

			special_ip, _ := special_list.NewService().Read(name)

			if special_ip != "" {
				resp := new(dns.Msg)
				resp.SetReply(r)
				resp.Id = r.Id // Set the response ID to match the query ID

				// Constructing an A record type (IPv4) response
				answer := &dns.A{
					Hdr: dns.RR_Header{Name: dns.Fqdn(name), Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 3600},
					A:   net.ParseIP(special_ip), // Convert the fetched IP string to net.IP
				}

				// Add the answer to the response
				resp.Answer = append(resp.Answer, answer)

				// Write the DNS message response
				w.WriteMsg(resp)
				endTime := time.Now()                            // 记录处理完成的时间
				elapsed := endTime.Sub(startTime).Milliseconds() // 计算处理耗时（毫秒）
				logger.Println("client: "+clientIP+" elapsedTime: ", elapsed, "ms  SepcialList Hit:  name: "+name+" ip: "+special_ip)
				return
			}

			/**
			查询缓存
			*/
			cachedResponse, found := lookupFromCache(queryKey)
			if found {
				resp := new(dns.Msg)
				resp.SetReply(r)
				resp.Answer = cachedResponse.Answer // Copy answers from cached response

				resp.Id = r.Id // Set the response ID to match the query ID

				w.WriteMsg(resp)
				endTime := time.Now()                            // 记录处理完成的时间
				elapsed := endTime.Sub(startTime).Milliseconds() // 计算处理耗时（毫秒）
				logger.Println("client: "+clientIP+" elapsedTime: ", elapsed, "ms  Cache hit for:", name)
				return
			}

			/**
			本地查询
			*/

			ip, err := net.LookupIP(name)
			if err == nil {
				for _, addr := range ip {
					ipStr := addr.String()
					if isPrivateIP(addr) {
						rr := &dns.A{
							Hdr: dns.RR_Header{Name: name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
							A:   addr,
						}
						m.Answer = append(m.Answer, rr)
						storeInCache(queryKey, m, appConfigs.DNS.CacheTimeout) // Cache the response
						w.WriteMsg(m)
						endTime := time.Now()                            // 记录处理完成的时间
						elapsed := endTime.Sub(startTime).Milliseconds() // 计算处理耗时（毫秒）
						logger.Println("client: "+clientIP+" elapsedTime: ", elapsed, "ms  name: ", name, " localPrivateDNS:", m.String())
						return
					}
					if isIPInRange(ipStr, ipRanges) {
						rr := &dns.A{
							Hdr: dns.RR_Header{Name: name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
							A:   addr,
						}
						m.Answer = append(m.Answer, rr)
						storeInCache(queryKey, m, appConfigs.DNS.CacheTimeout) // Cache the response
						w.WriteMsg(m)
						endTime := time.Now()                            // 记录处理完成的时间
						elapsed := endTime.Sub(startTime).Milliseconds() // 计算处理耗时（毫秒）
						logger.Println("client: "+clientIP+" elapsedTime: ", elapsed, "ms name: ", name, " localDNS:", m.String())
						return
					}
				}
			}

			/**
			递归查询
			*/
			recursiveAddr := appConfigs.DNS.RecursiveDNS
			c := new(dns.Client)
			resp, _, err := c.Exchange(r, recursiveAddr+":53")
			if err != nil {
				logger.Println("Error:", err)
				return
			}
			storeInCache(queryKey, resp, appConfigs.DNS.CacheTimeout) // Cache the response
			w.WriteMsg(resp)
			endTime := time.Now()                            // 记录处理完成的时间
			elapsed := endTime.Sub(startTime).Milliseconds() // 计算处理耗时（毫秒）
			logger.Println("client: "+clientIP+"  elapsedTime: ", elapsed, "ms name: ", name, "  recursiveDNS:", resp.String())
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
