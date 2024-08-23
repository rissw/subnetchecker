package main

import (
	"fmt"
	"net"
	"sort"
)

// Function to find the smallest subnet containing all given IPv4 addresses
func SmallestSubnetIPv4(ips []string) (*net.IPNet, error) {
	ipAddresses := make([]uint32, len(ips))
	for i, ip := range ips {
		parsedIP := net.ParseIP(ip).To4()
		if parsedIP == nil {
			return nil, fmt.Errorf("invalid IPv4 address: %s", ip)
		}
		ipAddresses[i] = ipToUint32(parsedIP)
	}

	// Sort IP addresses
	sort.Slice(ipAddresses, func(i, j int) bool {
		return ipAddresses[i] < ipAddresses[j]
	})

	// Find the smallest subnet that can contain both the smallest and largest IPs
	firstIP := ipAddresses[0]
	lastIP := ipAddresses[len(ipAddresses)-1]
	maskLen := commonPrefixLengthIPv4(firstIP, lastIP)

	// Create a subnet with the first IP and the mask
	mask := net.CIDRMask(maskLen, 32)
	subnetIP := uint32ToIP(firstIP & binaryMask(maskLen))

	return &net.IPNet{IP: subnetIP, Mask: mask}, nil
}

// Convert IPv4 address from net.IP to uint32
func ipToUint32(ip net.IP) uint32 {
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

// Convert uint32 to net.IP (IPv4)
func uint32ToIP(n uint32) net.IP {
	return net.IPv4(byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
}

// Find the common prefix length of two uint32 IP addresses
func commonPrefixLengthIPv4(ip1, ip2 uint32) int {
	diff := ip1 ^ ip2
	length := 0
	for diff > 0 {
		diff >>= 1
		length++
	}
	return 32 - length
}

// Generate binary mask from prefix length
func binaryMask(prefixLen int) uint32 {
	return ^uint32(0) << (32 - prefixLen)
}

func main() {
	// ips := []string{}
	ips := []string{"192.168.1.1", "192.168.1.10", "192.168.1.20"}
	subnet, err := SmallestSubnetIPv4(ips)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Smallest Subnet:", subnet.String())
}
