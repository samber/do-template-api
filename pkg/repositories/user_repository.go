package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
)

// User represents a user model
// This struct demonstrates how to define domain models for data access.
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository defines the interface for user data access operations
// This interface demonstrates how to define contracts for repository pattern.
type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context, limit, offset int) ([]*User, error)
}

// userRepository implements the UserRepository interface
// This struct demonstrates how to implement repository pattern with dependency injection.
type userRepository struct {
	db *pgxpool.Pool `do:""`
}

// NewUserRepository creates a new UserRepository instance
// This function demonstrates how to initialize a repository with database dependency.
func NewUserRepository(injector do.Injector) (UserRepository, error) {
	// Get database pool from the injector
	db := do.MustInvoke[*Database](injector)

	return &userRepository{db: db.Pool()}, nil
}

// CreateUser creates a new user in the database
// This method demonstrates how to implement CREATE operation with dependency injection.
func (r *userRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := `
		INSERT INTO users (name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, created_at, updated_at
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.CreatedAt, user.UpdatedAt).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
// This method demonstrates how to implement READ operation with dependency injection.
func (r *userRepository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return &user, nil
}

// GetUserByEmail retrieves a user by email
// This method demonstrates how to implement READ operation with dependency injection.
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &user, nil
}

// UpdateUser updates an existing user
// This method demonstrates how to implement UPDATE operation with dependency injection.
func (r *userRepository) UpdateUser(ctx context.Context, user *User) (*User, error) {
	query := `
		UPDATE users
		SET name = $1, email = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, name, email, created_at, updated_at
	`

	user.UpdatedAt = time.Now()

	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.UpdatedAt, user.ID).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

// DeleteUser deletes a user by ID
// This method demonstrates how to implement DELETE operation with dependency injection.
func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}

// ListUsers retrieves a list of users with pagination
// This method demonstrates how to implement LIST operation with dependency injection.
func (r *userRepository) ListUsers(ctx context.Context, limit, offset int) ([]*User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate users: %w", err)
	}

	return users, nil
}
