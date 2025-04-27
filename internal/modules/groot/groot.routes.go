package groot

import (
	"github.com/labstack/echo/v4"

	"go-api-boilerplate/internal/server/handler"
)

func RegisterGrootRoutes(e *echo.Echo, h *handler.Handler) {

	e.GET("/", HandleGroot(h))
}
