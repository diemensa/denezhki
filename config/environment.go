package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisHost string
	RedisPort string

	RabbitHost     string
	RabbitPort     string
	RabbitUser     string
	RabbitPassword string
	RabbitVHost    string
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env not found")
	}

	return &Config{
		Port: os.Getenv("PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),

		RabbitHost:     os.Getenv("RABBIT_HOST"),
		RabbitPort:     os.Getenv("RABBIT_PORT"),
		RabbitUser:     os.Getenv("RABBIT_USER"),
		RabbitPassword: os.Getenv("RABBIT_PASSWORD"),
		RabbitVHost:    os.Getenv("RABBIT_VHOST"),
	}
}
