package config

import "os"

// Config хранит все данные для конфигурации сервера
// которые беруться из переменных окружения
type Config struct {
	// MongoURI Хранит URI для подключения к базе данных
	MongoURI string

	// Port Порт на котором будет сервер
	Port string
}

func parseEnv(envName, defaultValue string) string {
	env := os.Getenv(envName)

	if env == "" {
		return defaultValue
	}

	return env
}

// New коструктор для конфига
func New() Config {
	return Config{
		MongoURI: parseEnv("MongoURI", "mongodb://localhost:27017"),
		Port:     parseEnv("Port", "4000"),
	}
}
