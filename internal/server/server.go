package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	"go-api-boilerplate/internal/database/queries"
	"go-api-boilerplate/internal/server/handler"
)

type Server struct {
	DB         *pgxpool.Pool
	Queries    *queries.Queries
	Logger     *slog.Logger
	HttpServer *http.Server
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

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080 // Default port
	}

	// Create the server instance
	s := &Server{
		DB:      pool,
		Queries: q,
		Logger:  logger,
	}
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      s.RegisterRoutes(h),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	s.HttpServer = httpServer

	return s
}

func (s *Server) Start() error {
	s.Logger.Info("Server starting up and listening on port:", "port", s.HttpServer.Addr)
	err := s.HttpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.Logger.Error("Error starting server:", "error", err)
	}
	return err
}
