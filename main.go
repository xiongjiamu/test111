package main

import (
	"encoding/json"
	"fmt"
	"github.com/miekg/dns"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type CachedResponse struct {
	Msg      *dns.Msg
	ExpireAt time.Time
}

type Config struct {
	RecursiveDNS    string        `json:"recursive_dns"`
	Whitelist       string        `json:"whitelist"`
	LogFile         string        `json:"log_file"`
	CacheTimeoutStr string        `json:"cache_timeout"`
	CacheTimeout    time.Duration // Add cache timeout configuration
}

var (
	logger     *log.Logger
	cache      map[string]CachedResponse
	cacheMutex sync.RWMutex
)

func loadConfig(filename string) (Config, error) {
	var config Config

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	// 转换字符串为 time.Duration
	if config.CacheTimeoutStr != "" {
		duration, err := time.ParseDuration(config.CacheTimeoutStr)
		if err != nil {
			return config, err
		}
		config.CacheTimeout = duration
	}
	return config, nil
}

// ... (Other functions remain unchanged)

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

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg, config Config) {
	m := new(dns.Msg)
	m.SetReply(r)

	ipRanges, err := loadIPRangesFromFile(config.Whitelist)
	if err != nil {
		logger.Println("Error loading IP ranges:", err)
		return
	}

	for _, q := range r.Question {
		name := strings.ToLower(q.Name)
		if q.Qtype == dns.TypeA {
			queryKey := fmt.Sprintf("%s-%d", name, q.Qtype)

			if cachedResponse, found := lookupFromCache(queryKey); found {
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
						storeInCache(queryKey, m, config.CacheTimeout) // Cache the response
						w.WriteMsg(m)
						return
					}
				}
			}

			recursiveAddr := config.RecursiveDNS
			c := new(dns.Client)
			resp, _, err := c.Exchange(r, recursiveAddr+":53")
			if err != nil {
				logger.Println("Error:", err)
				return
			}
			logger.Println("recursiveDNS:", resp.String())
			storeInCache(queryKey, resp, config.CacheTimeout) // Cache the response
			w.WriteMsg(resp)
		}
	}
}

func setupLogger(logFile string) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	logger = log.New(file, "DNS Server Log: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	cache = make(map[string]CachedResponse) // Initialize the cache

	setupLogger(config.LogFile)

	udpHandler := dns.NewServeMux()
	udpHandler.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		handleDNSRequest(w, r, config)
	})

	server := &dns.Server{Addr: ":53", Net: "udp", Handler: udpHandler}

	logger.Println("Starting DNS server at port 53...")
	if err := server.ListenAndServe(); err != nil {
		logger.Println("Error starting DNS server:", err)
		os.Exit(1)
	}
}
