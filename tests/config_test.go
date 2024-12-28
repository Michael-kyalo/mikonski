package tests

import (
	"os"
	"testing"

	"github.com/Michael-kyalo/mikonski/pkg/config"
)

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("OPENAI_API_KEY", "test-key")
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if cfg.APIKey != "test-key" {
		t.Errorf("Expected APIKey to be 'test-key', got %q", cfg.APIKey)
	}
	os.Unsetenv("OPENAI_API_KEY")
}

func TestLoadConfigFromFile(t *testing.T) {
	// Write a temporary .env file
	err := os.WriteFile(".env", []byte("OPENAI_API_KEY=test-file-key"), 0644)
	if err != nil {
		t.Fatalf("Failed to write .env file: %v", err)
	}
	defer os.Remove(".env")

	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if cfg.APIKey != "test-file-key" {
		t.Errorf("Expected APIKey to be 'test-file-key', got %q", cfg.APIKey)
	}
}

func TestCLIOverrides(t *testing.T) {
	// Simulate environment variables
	envConfig := config.Config{
		APIKey: "env-key",
		Model:  "env-model",
	}
	cliFlags := map[string]string{
		"APIKey": "cli-key",
		"Model":  "cli-model",
	}

	// Load configuration with overrides
	cfg := config.ApplyOverrides(&envConfig, cliFlags)

	// Verify that CLI flags take precedence
	if cfg.APIKey != "cli-key" {
		t.Errorf("Expected APIKey to be 'cli-key', got %q", cfg.APIKey)
	}
	if cfg.Model != "cli-model" {
		t.Errorf("Expected Model to be 'cli-model', got %q", cfg.Model)
	}
}
