package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/miekg/dns"
)

type Config struct {
	RecursiveDNS string `json:"recursive_dns"`
	Whitelist    string `json:"whitelist"`
	LogFile      string `json:"log_file"`
}

var logger *log.Logger

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

	return config, nil
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
