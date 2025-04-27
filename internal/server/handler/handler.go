package handler

import (
	"go-api-boilerplate/internal/database/queries"
)

// Handler is the single source of truth shared struct that contains dependencies needed by all handlers
type Handler struct {
	Q queries.Queries
}
