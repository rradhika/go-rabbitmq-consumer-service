package config

import (
	"os"
)

type Config struct {
	AppName string
	Port    string
}

func GetConfig() *Config {
	return &Config{
		AppName: getenv("APP_NAME", "GoFiberApp"),
		Port:    getenv("PORT", "3000"),
	}
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
