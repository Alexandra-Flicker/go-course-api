package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB struct {
		DSN string
	}
	Server struct {
		Port string
	}
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️ .env не найден, читаю переменные окружения напрямую")
	}

	cfg := &Config{}
	cfg.DB.DSN = os.Getenv("DB_DSN")
	cfg.Server.Port = os.Getenv("SERVER_PORT")

	return cfg
}
