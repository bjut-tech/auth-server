package middlewares

import (
	"github.com/bjut-tech/auth-server/app/utils"
	"github.com/bjut-tech/auth-server/internal/jwt"
	"github.com/labstack/echo/v4"
)

func BindCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := utils.ExtractToken(c.Request())

		token, err := jwt.ParseToken(tokenString)
		if err != nil {
			token = nil
		}

		var sub string
		if token != nil {
			sub, _ = token.Claims.GetSubject()
		}

		cc := &utils.CustomContext{
			Token:   token,
			Subject: sub,
			Context: c,
		}
		return next(cc)
	}
}
