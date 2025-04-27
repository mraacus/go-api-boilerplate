package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func Health(pool *pgxpool.Pool) map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := pool.Ping(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "Database connection healthy"

	// Get database stats
	dbStats := pool.Stat()
	stats["open_connections"] = strconv.Itoa(int(dbStats.TotalConns()))
	stats["aqcuired_connections"] = strconv.Itoa(int(dbStats.AcquiredConns()))
	stats["idle_connections"] = strconv.Itoa(int(dbStats.IdleConns()))
	stats["empty_acquire_count"] = strconv.Itoa(int(dbStats.EmptyAcquireCount()))
	stats["empty_acquire_wait_time_ms"] = strconv.Itoa(int(dbStats.EmptyAcquireWaitTime().Milliseconds()))
	stats["max_idle_destroy_count"] = strconv.Itoa(int(dbStats.MaxIdleDestroyCount()))
	stats["max_lifetime_destroy_count"] = strconv.Itoa(int(dbStats.MaxLifetimeDestroyCount()))

	// Evaluate stats to provide a health message
	if dbStats.AcquiredConns() > 40 { // For a max pool size of 50, this is 80% (Adjust as needed)
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.EmptyAcquireCount() > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleDestroyCount() > int64(dbStats.TotalConns())/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeDestroyCount() > int64(dbStats.TotalConns())/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}
