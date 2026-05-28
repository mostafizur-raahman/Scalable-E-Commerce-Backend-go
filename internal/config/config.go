package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var config Config

type Config struct {
	Port    int
	AppName string
	Version string
}

func loadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file")
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Failed to convert to int in .env file")
		os.Exit(1)
	}
	config = Config{
		Port:    port,
		AppName: os.Getenv("APP_NAME"),
		Version: os.Getenv("VERSION"),
	}
	return config

}

func Get() Config {
	return loadConfig()
}
