package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string
}

func Init() error {
	return godotenv.Load()
}

func GetServerConfig() ServerConfig {
	var config ServerConfig
	config.Port = os.Getenv("SERVER_PORT")
	return config
}
