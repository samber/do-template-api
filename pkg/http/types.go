package http

import (
	"time"
)

// CreateUserRequest represents the request body for creating a user
// This demonstrates how to define request DTOs with validation tags.
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// UpdateUserRequest represents the request body for updating a user
// This demonstrates how to define update request DTOs with validation.
type UpdateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// UserResponse represents the response body for user operations
// This demonstrates how to define response DTOs.
type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HealthResponse represents the response body for health checks.
type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}
