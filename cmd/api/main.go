package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"go-api-boilerplate/internal/server"
)

// gracefulShutdown handles graceful shutdown of the server when the user presses Ctrl+C via os signal context
func gracefulShutdown(s *server.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	s.Logger.Info("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := s.HttpServer.Shutdown(ctx); err != nil {
		s.Logger.Error("Server forced to shutdown with error", "error", err)
	}
	// Close the database connection
	s.DB.Close()

	s.Logger.Info("Server exiting")
	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	// Initialize the server instance
	server := server.NewServer()

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	// Start the server
	err := server.Start()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Wait for the graceful shutdown to complete
	<-done
	server.Logger.Info("Graceful shutdown complete.")
}
