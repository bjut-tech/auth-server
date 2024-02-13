package middlewares

import "github.com/labstack/echo/v4/middleware"

var csrfConfig = middleware.CSRFConfig{}

var CSRF = middleware.CSRFWithConfig(csrfConfig)
