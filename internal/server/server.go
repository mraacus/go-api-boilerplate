package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-api-boilerplate/internal/database/queries"
	"go-api-boilerplate/internal/server/handler"
	"go-api-boilerplate/internal/server/middlewares"
)

type Server struct {
	DB      *pgxpool.Pool
	Queries *queries.Queries
	Logger  *slog.Logger
	Echo    *echo.Echo
}

var ctx = context.Background()

var (
	database_name = os.Getenv("DB_DATABASE")
	password      = os.Getenv("DB_PASSWORD")
	username      = os.Getenv("DB_USER")
	port          = os.Getenv("DB_PORT")
	host          = os.Getenv("DB_HOST")
	schema        = os.Getenv("DB_SCHEMA")
)

func NewServer() *Server {
	// Set up logger using slog
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Initialize the database connection pool
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database_name, schema)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error("Failed to connect to database:", "error", err)
		os.Exit(1)
	}
	logger.Info("Successfully connected to database", "dsn", dsn)

	// Initialize sqlc queries with the database connection pool
	q := queries.New(pool)
	logger.Info("Successfully initialized sqlc queries")

	h := handler.Handler{
		Q:      *q,
		Logger: logger,
	}
	logger.Info("Successfully initialized handlers")

	// Create the server instance
	s := &Server{
		DB:      pool,
		Queries: q,
		Logger:  logger,
	}

	// Set up Echo with middleware and routes
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.CustomMiddleware)

	s.RegisterRoutes(e, h)

	s.Echo = e

	return s
}

func (s *Server) Start() error {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080 // Default port
	}
	s.Logger.Info("Server starting up and listening on port:", "port", port)

	// Configure server timeouts
	s.Echo.Server.ReadTimeout = 10 * time.Second
	s.Echo.Server.WriteTimeout = 30 * time.Second
	s.Echo.Server.IdleTimeout = time.Minute

	// Start the Echo server
	return s.Echo.Start(fmt.Sprintf(":%d", port))
}
