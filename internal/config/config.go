package config

import (
	"encoding/base64"
	"fmt"
	"log"
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

var env = os.Getenv("APP_ENV")

var CookieHost = os.Getenv("APP_COOKIE_HOST")
var CookieSecret = getEnvBytes("APP_COOKIE_SECRET")
var ListenAddr = "localhost:8080"
var Production = false

func init() {
	if env == "production" {
		Production = true
		ListenAddr = ":8080"

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
