package config

import (
	"flag"
	"os"
	"strconv"
	"time"
)

// Config хранит настройки для HTTP-сервера и таймаутов.
type Config struct {
	ListenHost   string
	ListenPort   int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// NewConfig создает и заполняет конфигурацию из переменных окружения и флагов командной строки.
func NewConfig() *Config {
	cfg := &Config{}

	// Установка значений по умолчанию
	cfg.ListenHost = getEnv("LISTEN_HOST", "0.0.0.0")
	cfg.ListenPort = getEnvAsInt("LISTEN_PORT", 8080)
	cfg.ReadTimeout = getEnvAsDuration("READ_TIMEOUT", 60*time.Second)
	cfg.WriteTimeout = getEnvAsDuration("WRITE_TIMEOUT", 60*time.Second)

	// Определение флагов командной строки
	flag.StringVar(&cfg.ListenHost, "host", cfg.ListenHost, "Host to listen on")
	flag.IntVar(&cfg.ListenPort, "port", cfg.ListenPort, "Port to listen on")
	flag.DurationVar(&cfg.ReadTimeout, "read-timeout", cfg.ReadTimeout, "HTTP server read timeout")
	flag.DurationVar(&cfg.WriteTimeout, "write-timeout", cfg.WriteTimeout, "HTTP server write timeout")

	flag.Parse() // Парсинг аргументов командной строки

	return cfg
}

// getEnv получает переменную окружения или значение по умолчанию.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt получает переменную окружения и возвращает ее как целое число.
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

// getEnvAsDuration получает переменную окружения и возвращает ее как объект time.Duration.
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
