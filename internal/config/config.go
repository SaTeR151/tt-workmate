package config

import (
	"os"
)

type ServerConfig struct {
	Port string
}

func GetServerConfig() ServerConfig {
	var config ServerConfig
	config.Port = os.Getenv("SERVER_PORT")
	return config
}
