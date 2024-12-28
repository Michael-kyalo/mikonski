package tests

import (
	"github.com/Michael-kyalo/mikonski/pkg/config"
	"os"
	"testing"
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
