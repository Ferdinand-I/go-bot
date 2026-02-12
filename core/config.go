package core

import (
	"fmt"
	"learning_bot/misc"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type BotConfig struct {
	Token string `envconfig:"BOT_TOKEN" required:"true"`
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     int    `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	DBName   string `envconfig:"DB_NAME" required:"true"`
}

type Config struct {
	BotConfig BotConfig
	DBConfig  DBConfig
}

func (cfg *DBConfig) BuildDSN() (string, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	return dsn, nil
}

func loadDotenvCfg() {
	misc.Must(godotenv.Load())
}

func buildCfg() Config {
	loadDotenvCfg()

	var BotCfg BotConfig
	var DBCfg DBConfig

	misc.Must(envconfig.Process("", &BotCfg))
	misc.Must(envconfig.Process("", &DBCfg))
	
	return Config{
		BotConfig: BotCfg,
		DBConfig:  DBCfg,
	}
}

var Cfg = buildCfg()
