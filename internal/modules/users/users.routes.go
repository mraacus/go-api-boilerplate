package users

import (
	"github.com/labstack/echo/v4"

	"go-api-boilerplate/internal/server/handler"
)

// RegisterUserRoutes registers all user-related routes
func RegisterUserRoutes(e *echo.Echo, h *handler.Handler) {

	e.POST("/users", func(c echo.Context) error {
		return CreateUser(c, h)
	})
	e.GET("/users", func(c echo.Context) error {
		return ListUsers(c, h)
	})
}
