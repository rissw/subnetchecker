package main

import (
	"fmt"
	"strconv"
	"time"

	"subnetchecker/pkg/sc"
)

func main() {
	/* netw, byte := sc.SubnetIPv4(
		net.IPv4(192, 168, 1, 65),
		net.IPv4(192, 168, 1, 120),
		net.IPv4(192, 168, 1, 123))
	fmt.Println(netw, byte)

	netw, byte = sc.MinMaxIP(
		net.IPv4(192, 168, 1, 65),
		net.IPv4(192, 168, 1, 120),
		net.IPv4(192, 168, 1, 123))
	fmt.Println(netw, byte)
	*/

	fmt.Println("Golang Running 100,000 loops")

	strIPs := []string{"192.168.1.192", "192.168.1.128", "192.168.1.133", "192.168.1.100"}
	t1 := time.Now()
	for i := 0; i < 100_000; i++ {
		_, _ = sc.MinMaxIP(strIPs)
	}
	t2 := time.Now()
	t3 := time.Now()
	for i := 0; i < 100_000; i++ {
		_, _ = sc.SubnetIPv4(strIPs)
	}
	t4 := time.Now()
	fmt.Println("Data", strIPs)
	fmt.Println("Golang Minmax    : ", float64(t2.UnixNano()-t1.UnixNano())/1_000_000_000)
	res, _ := sc.MinMaxIP(strIPs)
	fmt.Println("     Result      : ", res.String())
	fmt.Println("Golang Bitwise   : ", float64(t4.UnixNano()-t3.UnixNano())/1_000_000_000)
	res, _ = sc.SubnetIPv4(strIPs)
	fmt.Println("     Result      : ", res.String())

	strIPs = []string{
		"192.168.244.255",
		"172.31.255.255",
		"172.30.1.1",
		"192.168.1.1",
		"192.168.1.192",
		"192.168.1.128",
		"192.168.1.133",
		"192.168.1.100",
		"172.16.1.1",
		"128.1.1.90",
		"10.1.1.1",
		"10.2.2.2",
	}
	t1 = time.Now()
	for i := 0; i < 100_000; i++ {
		_, _ = sc.MinMaxIP(strIPs)
	}
	t2 = time.Now()
	t3 = time.Now()
	for i := 0; i < 100_000; i++ {
		_, _ = sc.SubnetIPv4(strIPs)
	}
	t4 = time.Now()
	fmt.Println()
	fmt.Println("Data", strIPs)
	fmt.Println("Golang Minmax    : ", float64(t2.UnixNano()-t1.UnixNano())/1_000_000_000)
	res, _ = sc.MinMaxIP(strIPs)
	fmt.Println("     Result      : ", res.String())
	fmt.Println("Golang Bitwise   : ", float64(t4.UnixNano()-t3.UnixNano())/1_000_000_000)
	res, _ = sc.SubnetIPv4(strIPs)
	fmt.Println("     Result      : ", res.String())

	strIPs = make([]string, 510)
	_prefix := "192.168.1."
	_prefix2 := "172.16.0."
	for i := 1; i < 256; i++ {
		strIPs[i-1] = _prefix + strconv.Itoa(i)
	}
	for i := 1; i < 256; i++ {
		strIPs[i+254] = _prefix2 + strconv.Itoa(i)
	}

	t1 = time.Now()
	for i := 0; i < 100_000; i++ {
		_, _ = sc.MinMaxIP(strIPs)
	}
	t2 = time.Now()
	t3 = time.Now()
	for i := 0; i < 100_000; i++ {
		_, _ = sc.SubnetIPv4(strIPs)
	}
	t4 = time.Now()
	fmt.Println()
	fmt.Println("Data 510 IPs")
	fmt.Println("Golang Minmax    : ", float64(t2.UnixNano()-t1.UnixNano())/1_000_000_000)
	res, _ = sc.MinMaxIP(strIPs)
	fmt.Println("     Result      : ", res.String())
	fmt.Println("Golang Bitwise   : ", float64(t4.UnixNano()-t3.UnixNano())/1_000_000_000)
	res, _ = sc.SubnetIPv4(strIPs)
	fmt.Println("     Result      : ", res.String())

}
