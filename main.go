package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/miekg/dns"
)

var whitelist = map[string]bool{
	"127.0.0.2": true, // 在白名单中的 IP 地址
	// 可以添加更多的 IP 地址到白名单中
}

func isIPInWhitelist(ip string) bool {
	_, found := whitelist[ip]
	return found
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)

	for _, q := range r.Question {
		name := strings.ToLower(q.Name)
		if q.Qtype == dns.TypeA {
			ip, err := net.LookupIP(name)
			if err == nil {
				for _, addr := range ip {
					ipStr := addr.String()
					if isIPInWhitelist(ipStr) {
						// 如果解析后的 IP 在白名单中，则直接返回该 IP
						rr := &dns.A{
							Hdr: dns.RR_Header{Name: name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
							A:   addr,
						}
						m.Answer = append(m.Answer, rr)
						w.WriteMsg(m)
						return
					}
				}
			}
			// 如果解析后的 IP 不在白名单中，则递归至 127.0.0.3 继续查询
			recursiveAddr := "127.0.0.3"
			c := new(dns.Client)
			resp, _, err := c.Exchange(r, recursiveAddr+":53")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			w.WriteMsg(resp)
		}
	}
}

func main() {
	udpHandler := dns.NewServeMux()
	udpHandler.HandleFunc(".", handleDNSRequest)

	server := &dns.Server{Addr: ":53", Net: "udp", Handler: udpHandler}

	fmt.Println("Starting DNS server at port 1053...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting DNS server:", err)
		os.Exit(1)
	}
}
