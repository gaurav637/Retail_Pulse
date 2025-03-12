package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	PORT       string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file.. not found env file: %v", err)
	}

	// Read DB_PORT, default to 3306 if empty
	portStr := os.Getenv("DB_PORT")
	if portStr == "" {
		portStr = "3306"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting DB_PORT to int: %v", err)
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     port,
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		PORT:       os.Getenv("SERVER_PORT"),
	}

	// Print config for debugging
	log.Printf("Loaded Config: %+v\n", config)

	return config
}
