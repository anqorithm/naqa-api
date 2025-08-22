package tests

import (
	"os"
	"testing"

	"github.com/anqorithm/naqa-api/internal/config"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed: %v", err)
	}

	if cfg.AppName == "" {
		t.Error("AppName should not be empty")
	}
	if cfg.Version == "" {
		t.Error("Version should not be empty")
	}
	if cfg.Port == "" {
		t.Error("Port should not be empty")
	}
}

func TestLoadConfigWithEnvVars(t *testing.T) {
	os.Setenv("APP_NAME", "Test App")
	os.Setenv("API_VERSION", "2.0.0")
	os.Setenv("PORT", "8080")
	defer func() {
		os.Unsetenv("APP_NAME")
		os.Unsetenv("API_VERSION")
		os.Unsetenv("PORT")
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed: %v", err)
	}

	if cfg.AppName != "Test App" {
		t.Errorf("Expected AppName 'Test App', got '%s'", cfg.AppName)
	}
	if cfg.Version != "2.0.0" {
		t.Errorf("Expected Version '2.0.0', got '%s'", cfg.Version)
	}
	if cfg.Port != "8080" {
		t.Errorf("Expected Port '8080', got '%s'", cfg.Port)
	}
}

func TestNewMongoConfig(t *testing.T) {
	cfg := config.NewMongoConfig()

	if cfg.URI == "" {
		t.Error("URI should not be empty")
	}
	if cfg.Database == "" {
		t.Error("Database should not be empty")
	}
	if cfg.Timeout == 0 {
		t.Error("Timeout should not be zero")
	}
}

func TestNewMongoConfigWithEnvVars(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://test:27017")
	os.Setenv("MONGO_DATABASE", "testdb")
	defer func() {
		os.Unsetenv("MONGO_URI")
		os.Unsetenv("MONGO_DATABASE")
	}()

	cfg := config.NewMongoConfig()

	if cfg.URI != "mongodb://test:27017" {
		t.Errorf("Expected URI 'mongodb://test:27017', got '%s'", cfg.URI)
	}
	if cfg.Database != "testdb" {
		t.Errorf("Expected Database 'testdb', got '%s'", cfg.Database)
	}
}