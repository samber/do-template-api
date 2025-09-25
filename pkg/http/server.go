package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/samber/do-template-api/pkg/config"
	"github.com/samber/do/v2"
)

// HTTPServer represents the HTTP server service
// This demonstrates how to create an HTTP server with dependency injection using do.
type HTTPServer struct {
	config        *config.Config `do:""`
	logger        zerolog.Logger `do:""`
	userHandler   *UserHandler   `do:""`
	healthHandler *HealthHandler `do:""`
	server        *http.Server
	engine        *gin.Engine
}

// NewHTTPServer creates a new HTTP server with dependency injection
// This function demonstrates how to initialize an HTTP server with all dependencies.
func NewHTTPServer(injector do.Injector) (*HTTPServer, error) {
	server := do.MustInvokeStruct[*HTTPServer](injector)

	// Setup Gin engine
	server.engine = gin.New()
	server.engine.Use(gin.Logger())
	server.engine.Use(gin.Recovery())

	// Setup routes
	server.setupRoutes()

	// Configure HTTP server using config from dependency injection
	server.server = &http.Server{
		Addr:         server.config.Server.Host + ":" + strconv.Itoa(server.config.Server.Port),
		Handler:      server.engine,
		ReadTimeout:  time.Duration(server.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(server.config.Server.WriteTimeout) * time.Second,
		IdleTimeout:  60 * time.Second, // Default idle timeout
	}

	return server, nil
}

// setupRoutes defines all the HTTP routes for the API
// This demonstrates how to organize routes in a structured way.
func (s *HTTPServer) setupRoutes() {
	api := s.engine.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	{
		users.POST("", s.userHandler.createUser)
		users.GET("", s.userHandler.listUsers)
		users.GET("/:id", s.userHandler.getUser)
		users.PUT("/:id", s.userHandler.updateUser)
		users.DELETE("/:id", s.userHandler.deleteUser)
	}

	// Health check endpoint
	s.engine.GET("/health", s.healthHandler.healthCheck)
}

// Start starts the HTTP server
// This demonstrates how to start the server with proper error handling.
func (s *HTTPServer) Start() error {
	s.logger.Info().
		Str("host", s.config.Server.Host).
		Int("port", s.config.Server.Port).
		Msg("Starting HTTP server")

	return s.server.ListenAndServe()
}

// Stop gracefully shuts down the HTTP server
// This demonstrates proper shutdown handling.
func (s *HTTPServer) ShutdownWithContext(ctx context.Context) error {
	s.logger.Info().Msg("Stopping HTTP server")
	return s.server.Shutdown(ctx)
}
