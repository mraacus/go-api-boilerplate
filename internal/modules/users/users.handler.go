package users

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"

	"go-api-boilerplate/internal/database/queries"
	"go-api-boilerplate/internal/server/handler"
)

func CreateUser(c echo.Context, h *handler.Handler) error {
	// Parse the request body
	req := new(CreateUserRequest)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}
	h.Logger.Info("Received a request to create a user", "request body", req)

	// Validate the request
	err = ValidateCreateUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	h.Logger.Info("Creating a user with", "name", req.Name, "role", req.Role)

	role := pgtype.Text{String: req.Role, Valid: true}
	// Create the user in the database
	user, err := h.Q.CreateUser(c.Request().Context(), queries.CreateUserParams{
		Name: req.Name,
		Role: role,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	h.Logger.Info("User created successfully", "user", user)

	response := map[string]any{
		"id":   user.ID,
		"name": user.Name,
		"role": user.Role.String,
	}

	return c.JSON(http.StatusOK, response)
}

func ListUsers(c echo.Context, h *handler.Handler) error {
	h.Logger.Info("Received a request to list users")
	users, err := h.Q.ListUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
