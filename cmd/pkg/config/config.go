package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port          string
	MapboxToken   string
	FirebaseCreds string
}

func Load() Config {
	_ = godotenv.Load() // Load .env file

	return Config{
		Port:          getEnv("PORT", "8080"),
		MapboxToken:   getEnv("MAPBOX_TOKEN", ""),
		FirebaseCreds: getEnv("FIREBASE_CREDS", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
