package main

import (
	"fmt"
	"net"
)

func main() {

	ip := "10.1.0.2"
	mask := "255.255.0.0"

	fmt.Printf("IP is :%s\n", net.ParseIP(ip))
	fmt.Printf("Mask is :%s\n", net.ParseIP(mask).DefaultMask())

	ipnet := &net.IPNet{
		IP:   net.ParseIP(ip),
		Mask: net.ParseIP(mask).DefaultMask(),
	}

	fmt.Println("IPNET IS:", ipnet)
}
