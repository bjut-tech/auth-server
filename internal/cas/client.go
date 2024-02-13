package cas

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

var canonicalBaseUrl = "https://cas.bjut.edu.cn"

var webvpnBaseUrl = "https://webvpn.bjut.edu.cn/https/77726476706e69737468656265737421f3f652d2253a7d44300d8db9d6562d"

var client = &http.Client{
	Timeout: 1 * time.Second,
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
	Timeout: 1 * time.Second,
}
