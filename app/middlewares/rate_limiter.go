package middlewares

import (
	"github.com/bjut-tech/auth-server/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"net/http"
	"strconv"
	"time"
)

func newRateLimiter() echo.MiddlewareFunc {
	rate := limiter.Rate{
		Period: 2 * time.Minute,
		Limit:  10,
	}
	store := memory.NewStore()
	ipRateLimiter := limiter.New(store, rate)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		if !config.Production {
			return next
		}

		return func(c echo.Context) (err error) {
			ip := c.RealIP()
			limiterCtx, err := ipRateLimiter.Get(c.Request().Context(), ip)
			if err != nil {
				return &echo.HTTPError{
					Code:     http.StatusInternalServerError,
					Internal: err,
				}
			}

			h := c.Response().Header()
			h.Set("Retry-After", strconv.FormatInt(limiterCtx.Reset-time.Now().Unix(), 10))

			if limiterCtx.Reached {
				return echo.NewHTTPError(http.StatusTooManyRequests, "请求频繁，请稍后再试")
			}

			return next(c)
		}
	}
}

var RateLimiter = newRateLimiter()
