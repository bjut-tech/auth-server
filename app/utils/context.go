package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	Token   *jwt.Token
	Subject string
	echo.Context
}
