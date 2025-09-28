package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// NewViper is a function to load config from config.json only
// No environment variable override support
func NewViper() *viper.Viper {
	config := viper.New()

	// Set config file
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")

	// Read config file
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Printf("Loaded config: %s\n", config.ConfigFileUsed())
	return config
}
