package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

type ApiConfig struct {
	ApiKey        string
	ApiIdentifier string
}
type Config struct {
	Database DBConfig
	API      ApiConfig
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("il faut un env")
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	return &Config{
		Database: DBConfig{
			DBUser: os.Getenv("DB_USER"),
			DBPass: os.Getenv("DB_PASS"),
			DBHost: os.Getenv("DB_HOST"),
			DBPort: port,
			DBName: os.Getenv("DB_NAME"),
		},
		API: ApiConfig{
			ApiKey:        os.Getenv("API_KEY"),
			ApiIdentifier: os.Getenv("API_ID_CLIENT"),
		},
	}
}
