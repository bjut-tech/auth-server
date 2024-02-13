package utils

import (
	"math/rand"
	"net"
)

func RandomIPv6() string {
	b := make([]byte, 16)

	b[0] = 0xfe
	b[1] = 0x80

	for i := 2; i < len(b); i++ {
		b[i] = byte(rand.Intn(256))
	}

	ipv6Addr := net.IP(b).String()
	return ipv6Addr
}
