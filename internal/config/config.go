package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP     HTTPConfig     `envPrefix:"HTTP_CONFIG"`
		Postgres PostgresConfig `envPrefix:"POSTGRES_URL"`
	}

	HTTPConfig struct {
		Port string `env:"HTTP_PORT" default:":8080"`
	}

	PostgresConfig struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT" default:"5432"`
		DB       string `env:"DB_NAME"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		SSLMode  string `env:"DB_SSLMODE" default:"disable"`
	}
)

var (
	config *Config
	once   sync.Once
)

func MustLoad() *Config {
	once.Do(func() {
		config = &Config{}
		if err := cleanenv.ReadEnv(config); err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	})

	return config
}

func (pgConfig *PostgresConfig) BuildConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.User,
		pgConfig.DB,
		pgConfig.Password,
		pgConfig.SSLMode,
	)
}
