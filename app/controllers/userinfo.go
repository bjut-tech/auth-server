package controllers

import (
	"github.com/bjut-tech/auth-server/app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUserInfo(c echo.Context) error {
	cc := c.(*utils.CustomContext)

	if cc.Token == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	return c.JSON(http.StatusOK, cc.Token.Claims)
}
