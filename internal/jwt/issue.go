package jwt

import (
	"github.com/bjut-tech/auth-server/internal/cas"
	"github.com/bjut-tech/auth-server/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtCustomClaims struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	jwt.RegisteredClaims
}

func CreateToken(user *cas.UserPrincipal, expiresIn time.Duration) string {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(expiresIn)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustomClaims{
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "https://auth.bjut.tech/",
			Subject:   user.Id,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})

	t, err := token.SignedString(config.CookieSecret)
	if err != nil {
		panic(err)
	}

	return t
}
