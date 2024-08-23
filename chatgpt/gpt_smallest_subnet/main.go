package main

import (
	"fmt"
	"net"
	"sort"
)

// Function to find the smallest subnet containing all given IPs
func SmallestSubnet(ips []string) (*net.IPNet, error) {
	// Convert IP strings to net.IP objects
	ipAddresses := make([]net.IP, len(ips))
	for i, ip := range ips {
		ipAddresses[i] = net.ParseIP(ip)
		if ipAddresses[i] == nil {
			return nil, fmt.Errorf("invalid IP address: %s", ip)
		}
	}

	// Sort IP addresses
	sort.Slice(ipAddresses, func(i, j int) bool {
		return compareIPs(ipAddresses[i], ipAddresses[j]) < 0
	})

	// Find the smallest subnet that can contain both the smallest and largest IPs
	firstIP := ipAddresses[0]
	lastIP := ipAddresses[len(ipAddresses)-1]
	maskLen := commonPrefixLength(firstIP, lastIP)

	// Create a subnet with the first IP and the mask
	mask := net.CIDRMask(maskLen, 8*len(firstIP))
	subnet := &net.IPNet{IP: firstIP.Mask(mask), Mask: mask}

	return subnet, nil
}

// Function to compare two IPs
func compareIPs(ip1, ip2 net.IP) int {
	for i := range ip1 {
		if ip1[i] < ip2[i] {
			return -1
		}
		if ip1[i] > ip2[i] {
			return 1
		}
	}
	return 0
}

// Function to find the common prefix length of two IPs
func commonPrefixLength(ip1, ip2 net.IP) int {
	length := 0
	for i := 0; i < len(ip1); i++ {
		for bit := 7; bit >= 0; bit-- {
			if (ip1[i]>>bit)&1 != (ip2[i]>>bit)&1 {
				return length
			}
			length++
		}
	}
	return length
}

func main() {
	ips := []string{"192.168.1.1", "192.168.1.10", "192.168.1.20"}
	subnet, err := SmallestSubnet(ips)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Smallest Subnet:", subnet.String())
}
