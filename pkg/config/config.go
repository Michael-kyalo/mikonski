package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

// Config holds the application configuration.
type Config struct {
	APIKey string
	Model  string
}

// LoadConfig loads the configuration from environment variables or a .env file.
func LoadConfig() (*Config, error) {
	// Load .env file if available
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("missing OPENAI_API_KEY")
	}

	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = "gpt-3.5-turbo" // Default model
	}

	return &Config{
		APIKey: apiKey,
		Model:  model,
	}, nil
}

// ApplyOverrides applies CLI flags to override loaded configuration.
func ApplyOverrides(cfg *Config, overrides map[string]string) *Config {
	if apiKey, exists := overrides["APIKey"]; exists {
		cfg.APIKey = apiKey
	}
	if model, exists := overrides["Model"]; exists {
		cfg.Model = model
	}
	return cfg
}
