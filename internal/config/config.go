package config

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func getEnvBytes(key string) []byte {
	val := os.Getenv(key)
	if strings.HasPrefix(val, "base64:") {
		val = strings.TrimPrefix(val, "base64:")
		decoded, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			panic(fmt.Errorf("error decoding %s: %w", key, err))
		}
		return decoded
	}
	return []byte(val)
}

func getEnvUrl(key string) *url.URL {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Errorf("%s is not set", key))
	}
	u, err := url.Parse(val)
	if err != nil {
		panic(fmt.Errorf("error parsing %s: %w", key, err))
	}
	return u
}

var env = os.Getenv("APP_ENV")

var BaseUrl = getEnvUrl("APP_BASE_URL")
var CookieHost = os.Getenv("APP_COOKIE_HOST")
var CookieSecret = getEnvBytes("APP_COOKIE_SECRET")
var ListenAddr = "localhost:8021"
var Production = false

func init() {
	if env == "production" {
		Production = true
		ListenAddr = ":8021"

		log.Println("Running in production mode")
	} else {
		log.Println("Running in development mode")
	}

	if CookieHost == "" {
		if Production {
			CookieHost = ".bjut.tech"
		} else {
			CookieHost = "localhost"
		}
	}

	if len(CookieSecret) < 32 {
		panic("cookie secret not set or too weak. use a key at least 256 bits long")
	}
}
