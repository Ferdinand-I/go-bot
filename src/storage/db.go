package storage

import (
	"learning_bot/core"
	"learning_bot/misc"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg core.DBConfig) *sqlx.DB {
	dsn := cfg.BuildDSN()

	db, err := sqlx.Connect("postgres", dsn)
	misc.Must(err)

	return db
}
