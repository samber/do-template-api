package cli

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/samber/do-template-api/pkg/config"
	httpservice "github.com/samber/do-template-api/pkg/http"
	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
)

// CLI represents the command line interface service
// This demonstrates how to create a CLI service with dependency injection.
type CLI struct {
	config      *config.Config `do:""`
	injector    do.Injector
	rootCommand *cobra.Command
}

// NewCLI creates a new CLI service with dependency injection support.
func NewCLI(i do.Injector) (*CLI, error) {
	cli := &CLI{
		config:   do.MustInvoke[*config.Config](i),
		injector: i,
	}

	// Create the root command
	cli.rootCommand = &cobra.Command{
		Use:     cli.config.App.Name,
		Short:   "A template api application using samber/do dependency injection",
		Long:    "A comprehensive template project demonstrating the github.com/samber/do dependency injection library with PostgreSQL and RabbitMQ integration",
		Version: cli.config.App.Version,
	}

	// Add persistent flags using dependency injection
	cli.setupPersistentFlags()

	// Add commands
	cli.setupCommands()

	return cli, nil
}

// setupPersistentFlags adds global flags to the CLI.
func (cli *CLI) setupPersistentFlags() {
	// Use the config service to set up all configuration flags
	// This demonstrates dependency injection for configuration management
	cli.config.SetCobraFlags(cli.rootCommand)
}

// setupCommands adds subcommands to the CLI.
func (cli *CLI) setupCommands() {
	// Add serve command
	cli.rootCommand.AddCommand(cli.newServeCommand())

	// Add migrate command
	cli.rootCommand.AddCommand(cli.newMigrateCommand())

	// Add health command
	cli.rootCommand.AddCommand(cli.newHealthCommand())

	// Add version command
	cli.rootCommand.AddCommand(cli.newVersionCommand())
}

// newServeCommand creates the serve command.
func (cli *CLI) newServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the api service",
		Long:  "Start the do-template-api service with dependency injection",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the HTTP server from the dependency injection container
			// This demonstrates how to access services from the CLI using do
			server := do.MustInvoke[*httpservice.HTTPServer](cli.injector)
			logger := do.MustInvoke[zerolog.Logger](cli.injector)

			// Setup graceful shutdown
			_, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Setup signal handling
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

			// Start server in goroutine
			go func() {
				logger.Info().Msg("Starting HTTP server...")
				if err := server.Start(); err != nil && err != http.ErrServerClosed {
					logger.Fatal().Err(err).Msg("Failed to start HTTP server")
				}
			}()
		},
	}
}

// newMigrateCommand creates the migrate command.
func (cli *CLI) newMigrateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		Long:  "Run database migrations using the configured database connection",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running database migrations...")
			// This will be implemented to use the dependency injection container
		},
	}
}

// newHealthCommand creates the health command.
func (cli *CLI) newHealthCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "health",
		Short: "Check service health",
		Long:  "Check the health of all services and dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Checking service health...")
			// This will be implemented to use the dependency injection container
		},
	}
}

// newVersionCommand creates the version command.
func (cli *CLI) newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Long:  "Show detailed version and build information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s version %s\n", cli.config.App.Name, cli.config.App.Version)
		},
	}
}

// RootCommand returns the root cobra command.
func (cli *CLI) RootCommand() *cobra.Command {
	return cli.rootCommand
}

// Execute executes the CLI with the given arguments.
func (cli *CLI) Execute() error {
	return cli.rootCommand.Execute()
}

// AddCommand adds a new command to the CLI.
func (cli *CLI) AddCommand(command *cobra.Command) {
	cli.rootCommand.AddCommand(command)
}
