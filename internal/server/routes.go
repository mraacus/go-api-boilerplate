package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-api-boilerplate/internal/database"
	"go-api-boilerplate/internal/modules/groot"
	"go-api-boilerplate/internal/modules/users"
	"go-api-boilerplate/internal/server/handler"
	"go-api-boilerplate/internal/server/middlewares"
)

func (s *Server) RegisterRoutes(handler handler.Handler) http.Handler {
	// Create a new Echo instance and set up middleware
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.CustomMiddleware)

	// Set up custom CORS middleware for the Echo instance
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Database health check route
	e.GET("/health", func(c echo.Context) error {
		healthData := database.Health(s.DB)
		return c.JSON(http.StatusOK, healthData)
	})

	// Register domain routes
	groot.RegisterGrootRoutes(e, &handler)
	users.RegisterUserRoutes(e, &handler)

	return e
}
