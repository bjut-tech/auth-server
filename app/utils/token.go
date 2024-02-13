package utils

import (
	"net/http"
	"strings"
)

func ExtractToken(req *http.Request) string {
	header := req.Header.Get("Authorization")
	if header != "" {
		return strings.TrimPrefix(header, "Bearer ")
	}

	cookie, _ := req.Cookie("_token")
	if cookie != nil {
		return cookie.Value
	}

	return ""
}
