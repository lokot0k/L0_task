package config

import (
	"os"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	KafkaBrokers     string
	AppIP            string
	AppPort          string
}

func LoadConfig() *Config {
	cfg := &Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		KafkaBrokers:     os.Getenv("KAFKA_BROKERS"),
		AppIP:            os.Getenv("APP_IP"),
		AppPort:          os.Getenv("APP_PORT"),
	}

	return cfg
}
