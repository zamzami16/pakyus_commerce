package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// NewViper is a function to load config based on environment
// Priority: ENV variables > config.{env}.json > config.json
func NewViper() *viper.Viper {
	config := viper.New()

	// Get environment (default: production)
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "production"
	}

	// Set config file name based on environment
	var configName string
	switch env {
	case "development", "dev":
		configName = "config.dev"
	case "testing", "test":
		configName = "config.test"
	default:
		configName = "config"
	}

	config.SetConfigName(configName)
	config.SetConfigType("json")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")

	// Try to read environment-specific config first
	err := config.ReadInConfig()
	if err != nil {
		// If environment-specific config not found, fallback to default config
		if env != "production" {
			config.SetConfigName("config")
			err = config.ReadInConfig()
			if err != nil {
				panic(fmt.Errorf("fatal error config file: %w", err))
			}
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	// Enable environment variable override
	config.AutomaticEnv()

	// Set environment variable prefixes for nested config
	config.SetEnvPrefix("PAKYUS")

	// Map environment variables to config keys
	config.BindEnv("database.username", "PAKYUS_DB_USERNAME")
	config.BindEnv("database.password", "PAKYUS_DB_PASSWORD")
	config.BindEnv("database.host", "PAKYUS_DB_HOST")
	config.BindEnv("database.port", "PAKYUS_DB_PORT")
	config.BindEnv("database.name", "PAKYUS_DB_NAME")
	config.BindEnv("web.port", "PAKYUS_WEB_PORT")
	config.BindEnv("kafka.bootstrap.servers", "PAKYUS_KAFKA_SERVERS")

	fmt.Printf("Loaded config: %s (env: %s)\n", config.ConfigFileUsed(), env)
	return config
}
