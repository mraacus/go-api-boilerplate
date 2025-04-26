package handlers

import (
	"net/http"

	"go-api-boilerplate/internal/database/queries"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Q queries.Queries
}

// groot mascot for root handler
func (h *Handler) GrootHandler(c echo.Context) error {
	// Can access h.DB here
	resp := map[string]string{
		"message": "I am groot",
	}
	return c.JSON(http.StatusOK, resp)
}

// func (h *Handler) HealthHandler(c echo.Context) error {
// 	return c.JSON(http.StatusOK, h.DB.Health())
// }

func (h *Handler) CreateUser(c echo.Context) error {
	// var user queries.User
	// if err := c.Bind(&user); err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }

	roleParam := pgtype.Text{String: "admin", Valid: true}

	user, err := h.Q.CreateUser(c.Request().Context(), queries.CreateUserParams{
		Name: "test user",
		Role: roleParam,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	response := map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
		"role": user.Role.String,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) ListUsers(c echo.Context) error {
	users, err := h.Q.ListUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
