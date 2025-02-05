package config

import (
	"os"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	// Установим переменные окружения для теста
	os.Setenv("LISTEN_HOST", "localhost")
	os.Setenv("LISTEN_PORT", "8080")
	os.Setenv("READ_TIMEOUT", "30s")
	os.Setenv("WRITE_TIMEOUT", "30s")

	cfg := NewConfig()

	if cfg.ListenHost != "localhost" {
		t.Errorf("Expected ListenHost to be 'localhost', got: %s", cfg.ListenHost)
	}
	if cfg.ListenPort != 8080 {
		t.Errorf("Expected ListenPort to be 8080, got: %d", cfg.ListenPort)
	}
	if cfg.ReadTimeout != 30*time.Second {
		t.Errorf("Expected ReadTimeout to be 30s, got: %v", cfg.ReadTimeout)
	}
	if cfg.WriteTimeout != 30*time.Second {
		t.Errorf("Expected WriteTimeout to be 30s, got: %v", cfg.WriteTimeout)
	}

	// Очистим переменные окружения
	os.Unsetenv("LISTEN_HOST")
	os.Unsetenv("LISTEN_PORT")
	os.Unsetenv("READ_TIMEOUT")
	os.Unsetenv("WRITE_TIMEOUT")
}
