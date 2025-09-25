package config

import (
	"fmt"
	"strings"

	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config holds all application configuration
// This struct demonstrates how to structure configuration for dependency injection
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Logger   LoggerConfig   `mapstructure:"logger"`
	App      AppConfig      `mapstructure:"app"`
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

// DatabaseConfig holds PostgreSQL configuration
type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Database        string `mapstructure:"database"`
	SSLMode         string `mapstructure:"ssl_mode"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

// LoggerConfig holds logger configuration
type LoggerConfig struct {
	Level   string `mapstructure:"level"`
	Format  string `mapstructure:"format"`
	Output  string `mapstructure:"output"`
	NoColor bool   `mapstructure:"no_color"`
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
	Debug       bool   `mapstructure:"debug"`
}

// NewConfig creates a new configuration instance using viper
// This demonstrates configuration management with the samber/do library
func NewConfig(i do.Injector) (*Config, error) {
	// Enable environment variable support
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	// Unmarshal configuration into struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

// SetCobraFlags adds command line flags to the cobra command
// This method demonstrates how services can provide functionality through DI
func (cs *Config) SetCobraFlags(cmd *cobra.Command) {
	// Server flags
	cmd.PersistentFlags().String("server.host", "localhost", "Server host")
	cmd.PersistentFlags().Int("server.port", 8080, "Server port")
	cmd.PersistentFlags().Int("server.read_timeout", 30, "Server read timeout in seconds")
	cmd.PersistentFlags().Int("server.write_timeout", 30, "Server write timeout in seconds")

	// Database flags
	cmd.PersistentFlags().String("database.host", "localhost", "Database host")
	cmd.PersistentFlags().Int("database.port", 5432, "Database port")
	cmd.PersistentFlags().String("database.user", "postgres", "Database user")
	cmd.PersistentFlags().String("database.password", "postgres", "Database password")
	cmd.PersistentFlags().String("database.database", "do_template_api", "Database name")
	cmd.PersistentFlags().String("database.ssl_mode", "disable", "Database SSL mode")
	cmd.PersistentFlags().Int("database.max_open_conns", 25, "Database max open connections")
	cmd.PersistentFlags().Int("database.max_idle_conns", 25, "Database max idle connections")
	cmd.PersistentFlags().Int("database.conn_max_lifetime", 300, "Database connection max lifetime in seconds")

	// Logger flags
	cmd.PersistentFlags().String("logger.level", "info", "Log level")
	cmd.PersistentFlags().String("logger.format", "console", "Log format")
	cmd.PersistentFlags().String("logger.output", "stdout", "Log output")
	cmd.PersistentFlags().Bool("logger.no_color", false, "Disable colored output")

	// App flags
	cmd.PersistentFlags().String("app.name", "do-template-worker", "Application name")
	cmd.PersistentFlags().String("app.version", "1.0.0", "Application version")
	cmd.PersistentFlags().String("app.environment", "development", "Application environment")
	cmd.PersistentFlags().Bool("app.debug", false, "Debug mode")

	// Bind all flags to viper for automatic configuration
	cs.bindFlagsToViper(cmd)
}

// bindFlagsToViper binds all cobra flags to viper
func (cs *Config) bindFlagsToViper(cmd *cobra.Command) {
	// Server flags
	viper.BindPFlag("server.host", cmd.PersistentFlags().Lookup("server.host"))
	viper.BindPFlag("server.port", cmd.PersistentFlags().Lookup("server.port"))
	viper.BindPFlag("server.read_timeout", cmd.PersistentFlags().Lookup("server.read_timeout"))
	viper.BindPFlag("server.write_timeout", cmd.PersistentFlags().Lookup("server.write_timeout"))

	// Database flags
	viper.BindPFlag("database.host", cmd.PersistentFlags().Lookup("database.host"))
	viper.BindPFlag("database.port", cmd.PersistentFlags().Lookup("database.port"))
	viper.BindPFlag("database.user", cmd.PersistentFlags().Lookup("database.user"))
	viper.BindPFlag("database.password", cmd.PersistentFlags().Lookup("database.password"))
	viper.BindPFlag("database.database", cmd.PersistentFlags().Lookup("database.database"))
	viper.BindPFlag("database.ssl_mode", cmd.PersistentFlags().Lookup("database.ssl_mode"))
	viper.BindPFlag("database.max_open_conns", cmd.PersistentFlags().Lookup("database.max_open_conns"))
	viper.BindPFlag("database.max_idle_conns", cmd.PersistentFlags().Lookup("database.max_idle_conns"))
	viper.BindPFlag("database.conn_max_lifetime", cmd.PersistentFlags().Lookup("database.conn_max_lifetime"))

	// Logger flags
	viper.BindPFlag("logger.level", cmd.PersistentFlags().Lookup("logger.level"))
	viper.BindPFlag("logger.format", cmd.PersistentFlags().Lookup("logger.format"))
	viper.BindPFlag("logger.output", cmd.PersistentFlags().Lookup("logger.output"))
	viper.BindPFlag("logger.no_color", cmd.PersistentFlags().Lookup("logger.no_color"))

	// App flags
	viper.BindPFlag("app.name", cmd.PersistentFlags().Lookup("app.name"))
	viper.BindPFlag("app.version", cmd.PersistentFlags().Lookup("app.version"))
	viper.BindPFlag("app.environment", cmd.PersistentFlags().Lookup("app.environment"))
	viper.BindPFlag("app.debug", cmd.PersistentFlags().Lookup("app.debug"))
}
