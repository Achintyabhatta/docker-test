package main

import (
	"net"
)

// helper function that performs an IPLookup if IP not already available in the database
func searchIP(domain string) string {
	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}
