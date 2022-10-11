package main

import (
	"log"
	"net"
)

func main() {
	host := "www.google.com"
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Println(err)
		return
	}
	for _, ip := range ips {
		log.Printf("%s -> %s\n", host, ip.String())
	}
}
