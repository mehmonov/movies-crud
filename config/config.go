package config

import (
    "os"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    ServerPort string
    JWTSecret  string
    AccessTokenSecret  string
    RefreshTokenSecret string
}

func NewConfig() *Config {
    return &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "1234"),
        DBName:     getEnv("DB_NAME", "movies-crud"),
        ServerPort: getEnv("SERVER_PORT", "8080"),
        JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
        AccessTokenSecret:  getEnv("ACCESS_TOKEN_SECRET", "your-access-token-secret"),
        RefreshTokenSecret: getEnv("REFRESH_TOKEN_SECRET", "your-refresh-token-secret"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}