package jwt

import (
	"fmt"
	"github.com/bjut-tech/auth-server/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return config.CookieSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, &ErrInvalidToken{}
	}

	_, err = token.Claims.GetSubject()
	if err != nil {
		return nil, &ErrInvalidToken{}
	}

	return token, nil
}
