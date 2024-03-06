package cas

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var canonicalBaseUrl = "https://cas.bjut.edu.cn"

var client = &http.Client{
	Timeout: 2 * time.Second,
}

var clientPinned = &http.Client{
	Transport: &http.Transport{
		DialTLSContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
			c, err := tls.Dial("tcp", "bjutwaf.bjut.tech:443", &tls.Config{
				ServerName: "cas.bjut.edu.cn",
			})
			if err != nil {
				return nil, err
			}
			return c, c.HandshakeContext(ctx)
		},
	},
	Timeout: 2 * time.Second,
}
