package controllers

import (
	"github.com/bjut-tech/auth-server/internal/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "_token",
		Value:    "",
		Domain:   config.CookieHost,
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return c.NoContent(http.StatusNoContent)
}
