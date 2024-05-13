package config

import (
	"os"
)

type Config struct {
	TelegramToken string
	WeatherToken  string
}

// GetConfig возвращяет структуру конфиг
func GetConfig() *Config {
	return &Config{
		TelegramToken: getEnv("TELEGRAM_TOKEN", ""),
		WeatherToken:  getEnv("WEATHER_TOKEN", ""),
	}
}

// если значение не будет найдено, то вернется дефолтное значение
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
