package config

import (
    "os"
)

type Config struct {
    DatabaseDSN string
    Port        string
}

func LoadConfig() *Config {
    return &Config{
        DatabaseDSN: getEnv("DATABASE_DSN", "file:test.db"),
        Port:        getEnv("PORT", ":8080"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
