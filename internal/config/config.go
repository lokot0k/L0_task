package config

import (
	"os"
	"strconv"
)

type Config struct {
	PostgresUser       string
	PostgresPassword   string
	PostgresDB         string
	PostgresHost       string
	PostgresPort       string
	KafkaBrokers       string
	AppIP              string
	AppPort            string
	KafkaTopic         string
	KafkaConsumerGroup string
	CacheLimit         int
}

func LoadConfig() *Config {
	cacheLimit, err := strconv.Atoi(os.Getenv("CACHE_LIMIT"))
	if err != nil {
		panic("Cache limit is not properly defined!")
	}
	cfg := &Config{
		PostgresUser:       os.Getenv("POSTGRES_USER"),
		PostgresPassword:   os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:         os.Getenv("POSTGRES_DB"),
		PostgresHost:       os.Getenv("POSTGRES_HOST"),
		PostgresPort:       os.Getenv("POSTGRES_PORT"),
		KafkaBrokers:       os.Getenv("KAFKA_BROKERS"),
		AppIP:              os.Getenv("APP_IP"),
		AppPort:            os.Getenv("APP_PORT"),
		KafkaTopic:         os.Getenv("KAFKA_TOPIC"),
		KafkaConsumerGroup: os.Getenv("KAFKA_CONSUMER_GROUP"),
		CacheLimit:         cacheLimit,
	}

	return cfg
}
