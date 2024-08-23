package sc

import (
	"fmt"
	"net"
)

func SubnetIPv4(strIPs []string) (*net.IPNet, error) {
	dataLength := len(strIPs)
	if dataLength == 0 {
		return &net.IPNet{
			IP:   net.IPv4(0, 0, 0, 0),
			Mask: []byte{0, 0, 0, 0},
		}, nil
	}
	ips := make([]uint32, dataLength)
	for i, strIP := range strIPs {
		ip := net.ParseIP(strIP).To4()
		if ip == nil {
			return nil, fmt.Errorf("invalid IP: %s", strIP)
		}
		ips[i] = ipv4ToUint32(ip)
	}

	// ---------
	netw := ips[0]
	mask := uint32(0xffffffff) // ^(netw ^ ipv4ToUint32(ips[1]))
	netw = netw & mask
	for i := 1; i < dataLength; i++ {
		mask &= ^(netw ^ ips[i])
		netw = netw & mask
	}
	mb := maskbit(mask)
	netw &= maskbitToUint32(mb)

	return &net.IPNet{
		IP:   net.IPv4(byte(netw>>24), byte(netw>>16), byte(netw>>8), byte(netw)),
		Mask: net.CIDRMask(int(mb), 32),
	}, nil
}

func ipv4ToUint32(ip net.IP) uint32 {
	ip4 := ip.To4()
	if ip4 == nil {
		return 0
	}
	return uint32(ip4[0])<<24 | uint32(ip4[1])<<16 | uint32(ip4[2])<<8 | uint32(ip4[3])
}

func MinMaxIP(strIPs []string) (*net.IPNet, error) {
	dataLength := len(strIPs)
	if dataLength == 0 {
		return &net.IPNet{
			IP:   net.IPv4(0, 0, 0, 0),
			Mask: []byte{0, 0, 0, 0},
		}, nil
	}
	ips := make([]uint32, dataLength)
	for i, strIP := range strIPs {
		ip := net.ParseIP(strIP).To4()
		if ip == nil {
			return nil, fmt.Errorf("invalid IP: %s", strIP)
		}
		ips[i] = ipv4ToUint32(ip)
	}
	min, max := ips[0], ips[0]
	for _, ip := range ips {
		if ip < min {
			min = ip
			continue
		}
		if ip > max {
			max = ip
		}
	}
	mb := maskbit(^(max ^ min))
	min &= maskbitToUint32(mb)
	return &net.IPNet{
		IP:   net.IPv4(byte(min>>24), byte(min>>16), byte(min>>8), byte(min)),
		Mask: net.CIDRMask(int(mb), 32),
	}, nil
}

func toIPv4(u uint32) net.IP {
	return net.IPv4(byte(u>>24), byte(u>>16), byte(u>>8), byte(u))
}

func maskbit(x uint32) (mb byte) {
	for i := 0; i < 4; i++ {
		// lk := lookupMask[x>>24]
		// mb += lk
		// if lk != 8 {
		// 	return
		// }

		b := x >> 24
		mb += lookupMask[b]
		if lookupMask[b] != 8 {
			return
		}
		x <<= 8
	}
	return
}

func maskbitToUint32(m byte) uint32 {
	return 0xffffffff << (32 - m)
}

const lookupMask = "" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01" +
	"\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01" +
	"\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01" +
	"\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01" +
	"\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02" +
	"\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02\x02" +
	"\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03\x03" +
	"\x04\x04\x04\x04\x04\x04\x04\x04\x05\x05\x05\x05\x06\x06\x07\x08"
