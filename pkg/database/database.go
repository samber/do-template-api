package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do-template-api/pkg/config"
	"github.com/samber/do/v2"
)

// Database represents a PostgreSQL connection pool
// This service demonstrates how to create and manage database connections using dependency injection.
type Database struct {
	pool *pgxpool.Pool
}

// NewDatabase creates a new PostgreSQL database connection pool
// This function demonstrates how to initialize a service with dependencies using samber/do.
func NewDatabase(injector do.Injector) (*Database, error) {
	// Get configuration from the injector
	appConfig := do.MustInvoke[*config.Config](injector)
	cfg := appConfig.Database

	// Build connection string
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s pool_max_conns=%d pool_min_conns=%d pool_max_conn_lifetime=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
		cfg.MaxOpenConns,
		cfg.MaxIdleConns,
		time.Duration(cfg.ConnMaxLifetime)*time.Second,
	)

	// Create connection pool config
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	// Additional pool configuration
	poolConfig.MaxConns = int32(cfg.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.MaxIdleConns)
	poolConfig.MaxConnLifetime = time.Duration(cfg.ConnMaxLifetime) * time.Second
	poolConfig.HealthCheckPeriod = 1 * time.Minute
	poolConfig.MaxConnIdleTime = 5 * time.Minute

	// Create connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{pool: pool}, nil
}

// Pool returns the underlying pgxpool.Pool
// This method demonstrates how to expose dependencies to other services.
func (db *Database) Pool() *pgxpool.Pool {
	return db.pool
}

// Health checks the database connection
// This method demonstrates how to implement health checks for services.
func (db *Database) HealthCheckWithContext(ctx context.Context) error {
	if err := db.pool.Ping(ctx); err != nil {
		return fmt.Errorf("database health check failed: %w", err)
	}
	return nil
}

func (db *Database) Shutdown(context.Context) error {
	if db.pool != nil {
		db.pool.Close()
	}

	return nil
}
