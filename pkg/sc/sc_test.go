package sc

import (
	"net"
	"testing"
)

var testData = []net.IP{
	{192, 168, 1, 1},
	{192, 168, 1, 12},
	{192, 168, 1, 16},
	{192, 168, 1, 130},
	{192, 168, 1, 111},
	{192, 168, 1, 221},
	{192, 168, 1, 14},
	{192, 168, 1, 19},
	{192, 168, 1, 192},
	{192, 168, 1, 61},
	{192, 168, 1, 119},
	{192, 168, 1, 192},
	{172, 1, 1, 1},
	{10, 1, 1, 1},
}

func BenchmarkMinMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinMaxIP(testData...)
	}
}

func BenchmarkSubnetIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubnetIPv4(testData...)
	}
}
