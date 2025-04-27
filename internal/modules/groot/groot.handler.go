package groot

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-api-boilerplate/internal/server/handler"
)

// groot mascot for groot handler (for testing)
func GrootHandler(c echo.Context, h *handler.Handler) error {
	h.Logger.Info("Received a request to groot")
	resp := map[string]string{
		"message": "I am groot",
	}
	return c.JSON(http.StatusOK, resp)
}
