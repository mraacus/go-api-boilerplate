package handlers

import (
	"net/http"

	"go-api-boilerplate/internal/database"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB database.Service
}

// groot mascot for root handler
func (h *Handler) GrootHandler(c echo.Context) error {
	// Can access h.DB here
	resp := map[string]string{
		"message": "I am groot",
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) HealthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, h.DB.Health())
}
