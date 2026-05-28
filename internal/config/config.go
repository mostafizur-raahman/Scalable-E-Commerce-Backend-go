package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             int
	AppName          string
	Version          string
	DatabaseURL      string
	DatabaseMaxConns int32
}

var (
	config Config
	once   sync.Once
)

func loadConfig() {
	_ = godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8000
	}

	maxConns, err := strconv.Atoi(os.Getenv("DB_MAX_CONNS"))
	if err != nil {
		maxConns = 10
	}

	config = Config{
		Port:             port,
		AppName:          getEnv("APP_NAME", "E-Commerce API"),
		Version:          getEnv("VERSION", "1.0.0"),
		DatabaseURL:      getEnv("DB_URL", "postgres://localhost:5432/ecommerce?sslmode=disable"),
		DatabaseMaxConns: int32(maxConns),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func Get() Config {
	once.Do(loadConfig)
	return config
}
