package internal

import (
	"errors"
	"github.com/bjut-tech/auth-server/app"
	"github.com/bjut-tech/auth-server/internal/config"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func StartServer() {
	e := echo.New()
	e.Debug = !config.Production
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.Validator = &customValidator{}

	app.RegisterRoutes(e)

	if err := e.Start(config.ListenAddr); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
