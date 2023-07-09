package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver   string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
}

// init - loads values from .env into the system
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// GetConfig - Returns a populated Config object with database connection configuration based on environment variables or default values.
func GetConfig() Config {
	return Config{
		DBDriver:   getEnv("DB_DRIVER", "postgres"),
		DBUser:     getEnv("DB_USER", "bestwallet"),
		DBPassword: getEnv("DB_PASSWORD", "bestwallet"),
		DBName:     getEnv("DB_NAME", "bestwallet"),
		DBHost:     getEnv("DB_HOST", "localhost"),
	}
}

// GetDBConnectionString - defined on the Config type, returns a formatted database connection string based on the configuration values. It also logs the connection string before returning it.
func (c Config) GetDBConnectionString() string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBUser, c.DBPassword, c.DBName)
	log.Println("Will connect to the following database: ", connectionString)
	return connectionString
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
