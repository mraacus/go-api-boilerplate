package middlewares

import (
	"github.com/labstack/echo/v4"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	// Add custom middleware logic here
	return func(c echo.Context) error {
		return next(c)
	}
}
