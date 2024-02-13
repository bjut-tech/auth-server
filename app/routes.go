package app

import (
	"github.com/bjut-tech/auth-server/app/controllers"
	"github.com/bjut-tech/auth-server/app/middlewares"
	"github.com/bjut-tech/auth-server/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func RegisterRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middlewares.BindCustomContext)
	e.Use(middleware.Recover())

	e.POST("/login", controllers.PostLogin, middlewares.RateLimiter)

	e.GET("/logout", controllers.Logout)
	e.POST("/logout", controllers.Logout)

	e.GET("/authorize", controllers.Authorize)
	e.POST("/authorize", controllers.Authorize)

	e.GET("/userinfo", controllers.GetUserInfo)

	e.GET("/*", echo.WrapHandler(http.FileServer(web.HttpFs())))
}
