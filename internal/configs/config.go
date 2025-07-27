package configs

import "os"

type Config struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
	SSLMode  string `env:"SSL_MODE"`
}

func New() *Config {
	return &Config{
		Host:     getEnv("HOST", ""),
		Port:     getEnv("PORT", ""),
		User:     getEnv("USER", ""),
		Password: getEnv("PASSWORD", ""),
		Database: getEnv("DATABASE", ""),
		SSLMode:  getEnv("SSL_MODE", "disable"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
