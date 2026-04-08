package config

import "os"

const (
	defaultHost        = "0.0.0.0"
	defaultPort        = "8080"
	defaultGreenAPIURL = "https://api.green-api.com"
)

type Config struct {
	Host        string
	Port        string
	GreenAPIURL string
}

func Load() Config {
	return Config{
		Host:        getenv("HOST", defaultHost),
		Port:        getenv("PORT", defaultPort),
		GreenAPIURL: getenv("GREEN_API_URL", defaultGreenAPIURL),
	}
}

func (c Config) Addr() string {
	return c.Host + ":" + c.Port
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
