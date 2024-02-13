package controllers

import (
	"errors"
	"github.com/bjut-tech/auth-server/internal/cas"
	"github.com/bjut-tech/auth-server/internal/config"
	"github.com/bjut-tech/auth-server/internal/jwt"
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type loginForm struct {
	Username string `form:"username" json:"username" validate:"required" label:"用户名"`
	Password string `form:"password" json:"password" validate:"required" label:"密码"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func PostLogin(c echo.Context) error {
	var form loginForm
	if err := c.Bind(&form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "请求无效")
	}
	if err := c.Validate(&form); err != nil {
		message := "请求无效"
		var validationErrors validate.Errors
		if errors.As(err, &validationErrors) {
			message = validationErrors.Random()
		}
		return echo.NewHTTPError(http.StatusUnprocessableEntity, message)
	}

	user, err := cas.GetUser(c.Request().Context(), form.Username, form.Password)
	if err != nil {
		if errors.Is(err, &cas.ErrInvalidCredentials{}) {
			return echo.NewHTTPError(http.StatusUnauthorized, "用户名或密码错误")
		} else if errors.Is(err, &cas.ErrThrottled{}) {
			return echo.NewHTTPError(http.StatusTooManyRequests, "请求过于频繁，请稍后再试")
		}
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  "服务器遇到错误，无法完成统一认证",
			Internal: err,
		}
	}

	expiresIn := time.Duration(24) * time.Hour
	token := jwt.CreateToken(user, expiresIn)
	c.SetCookie(&http.Cookie{
		Name:     "_token",
		Value:    token,
		Domain:   config.CookieHost,
		MaxAge:   int(expiresIn.Seconds()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return c.JSON(http.StatusOK, &tokenResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int(expiresIn.Seconds()),
	})
}
