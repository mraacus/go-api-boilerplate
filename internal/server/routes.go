package server

import (
	"go-api-boilerplate/internal/server/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes(handler handlers.Handler) http.Handler {
	// Create a new Echo instance and set up middleware
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set up custom CORS middleware for the Echo instance
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.GET("/", handler.GrootHandler)
	e.POST("/users", handler.CreateUser)
	// e.GET("/health", handler.HealthHandler)

	return e
}
