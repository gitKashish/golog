package config

import (
	"os"
	"strconv"
)

// Config holds all the configuration for the application
type Config struct {
	// Server configuration
	Server ServerConfig
	// Template configuration
	Template TemplateConfig
}

type ServerConfig struct {
	// Port to run the server on
	Port int
}

type TemplateConfig struct {
	// Path to the template file
	TemplatePath string
}

// NewConfig creates a new configuration with default values
// and overrides from environment variables
func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnvInt("PORT", 2600),
		},
		Template: TemplateConfig{
			TemplatePath: getEnvString("GOLOG_TEMPLATE_PATH", "template.yaml"),
		},
	}
}

// Helper function to get an environment variable as an integer
func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// Helper function to get an environment variable as a string
func getEnvString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
