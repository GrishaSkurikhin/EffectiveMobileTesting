package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	Rest        restServer
	Graphql     graphQLServer
	Cache       cache
	UserStorage userStorage
}

type restServer struct {
	Address     string
	Timeout     time.Duration
	IdleTimeout time.Duration
}

type graphQLServer struct {
	Address     string
}

type cache struct {
}

type userStorage struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func MustLoad(configPath string) {
	if configPath == "" {
		log.Fatal("config path is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	if err := godotenv.Load(configPath); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}
}

func New() *Config {
	return &Config{
		Env: getEnvString("ENV", "local"),
		Rest: restServer{
			Address:     getEnvString("REST_ADDRESS", "localhost:8080"),
			Timeout:     getEnvTime("REST_TIMEOUT", 4*time.Second),
			IdleTimeout: getEnvTime("REST_IDLE_TIMEOUT", 4*time.Second),
		},
		Graphql: graphQLServer{
			Address: getEnvString("GRAPHQL_ADDRESS", "localhost:4000"),
		},
		Cache: cache{},
		UserStorage: userStorage{
			Host:     getEnvString("USER_STORAGE_HOST", "localhost"),
			Port:     getEnvString("USER_STORAGE_PORT", "5432"),
			User:     getEnvString("USER_STORAGE_USER", "postgres"),
			Password: getEnvString("USER_STORAGE_PASSWORD", ""),
			Name:     getEnvString("USER_STORAGE_NAME", ""),
		},
	}
}

func getEnvString(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvTime(name string, defaultVal time.Duration) time.Duration {
	valueStr := getEnvString(name, "")
	if value, err := time.ParseDuration(valueStr); err == nil && valueStr != "" {
		return value
	}
	return defaultVal
}
