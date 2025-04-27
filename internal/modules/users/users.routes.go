package users

import (
	"github.com/labstack/echo/v4"

	"go-api-boilerplate/internal/server/handler"
)

func RegisterUserRoutes(e *echo.Echo, h *handler.Handler) {

	e.POST("/users", HandleCreateUser(h))
	e.GET("/users", HandleListUsers(h))
}
