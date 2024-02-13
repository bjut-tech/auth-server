package controllers

import (
	"github.com/bjut-tech/auth-server/app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"slices"
)

type authorizeParams struct {
	Redirect     string   `query:"redirect" json:"redirect"`
	AllowedUsers []string `query:"allowed" json:"allowed_users"`
}

func Authorize(c echo.Context) error {
	var params authorizeParams
	_ = c.Bind(&params) // params are optional, so errors are ignored

	cc := c.(*utils.CustomContext)

	if cc.Token == nil {
		if params.Redirect != "" {
			redirectUrl := c.Request().URL
			redirectUrl.Path = "/login.html"
			redirectUrl.RawQuery = "redirect=" + url.QueryEscape(params.Redirect)
			return c.Redirect(http.StatusFound, redirectUrl.String())
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("X-User-ID", cc.Subject)

	if params.AllowedUsers != nil && !slices.Contains(params.AllowedUsers, cc.Subject) {
		if params.Redirect != "" {
			redirectUrl := c.Request().URL
			redirectUrl.Path = "/index.html"
			redirectUrl.RawQuery = "unauthorized"
			return c.Redirect(http.StatusFound, redirectUrl.String())
		}

		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to access this resource")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"error":   nil,
		"message": "Authorized",
	})
}
