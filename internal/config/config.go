package config

import (
	"os"
	"strconv"
)

type Config struct {
	DataStore DataStoreConfig
	Broker    BrokerConfig
}

type DataStoreConfig struct {
	DataStoreType string

	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

type BrokerConfig struct {
	BrokerType string
	Host       string
	Port       int
	Username   string
	Password   string
}

func LoadConfig() *Config {
	dbConfig := DataStoreConfig{
		DataStoreType: getEnvWithDefault("ESR_DATASTORE_TYPE", "bbolt"),
		Host:          getEnv("ESR_DATASTORE_HOST"),
		Port:          getIntEnv("EST_DATASTORE_PORT"),
		Username:      getEnv("ESR_DATASTORE_USERNAME"),
		Password:      getEnv("ESR_DATASTORE_PASSWORD"),
		Name:          getEnvWithDefault("ESR_DATABASE_NAME", "esrdb"),
	}

	brokerConfig := BrokerConfig{
		BrokerType: getEnvWithDefault("ESR_BROKER_TYPE", "rabbitmq"),
		Host:       getEnvWithDefault("ESR_BROKER_HOST", "localhost"),
		Port:       getIntEnvWithDefault("ESR_BROKER_PORT", 5672),
		Username:   getEnvWithDefault("ESR_BROKER_USERNAME", "guest"),
		Password:   getEnvWithDefault("ESR_BROKER_PASSWORD", "guest"),
	}

	return &Config{
		DataStore: dbConfig,
		Broker:    brokerConfig,
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func getEnvWithDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getIntEnv(key string) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return i
}

func getIntEnvWithDefault(key string, fallback int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return i
}
