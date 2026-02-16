package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Env int

const (
	Local Env = iota
	Development
	Production
)

var envMap = map[string]Env{
	"LOCAL":       Local,
	"DEVELOPMENT": Development,
	"PRODUCTION":  Production,
}

type Bot struct {
	Token string `envconfig:"BOT_TOKEN" required:"true"`
}

type DB struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     int    `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	DBName   string `envconfig:"DB_NAME" required:"true"`
}

type Config struct {
	Bot Bot
	DB  DB
}

func (cfg *DB) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}

func Load() (*Config, error) {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		return nil, fmt.Errorf("ENVIRONMENT is not set")
	}

	if envMap[env] == Local {
		if err := godotenv.Load(".env"); err != nil {
			return nil, fmt.Errorf("loading .env: %w", err)
		}
	}

	var botCfg Bot
	if err := envconfig.Process("", &botCfg); err != nil {
		return nil, fmt.Errorf("processing bot config: %w", err)
	}

	var dbCfg DB
	if err := envconfig.Process("", &dbCfg); err != nil {
		return nil, fmt.Errorf("processing db config: %w", err)
	}

	return &Config{
		Bot: botCfg,
		DB:  dbCfg,
	}, nil
}
