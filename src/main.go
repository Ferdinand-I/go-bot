package main

import (
	"context"
	"learning_bot/bot"
	"learning_bot/core"
	"learning_bot/misc"
	"learning_bot/storage"
)

func main() {
	db := storage.ConnectDB(core.Cfg.DBConfig)
	defer db.Close()

	ctx := context.Context(misc.NewContext(context.Background(), db))
	b := bot.SetupBot(ctx)

	b.Start(ctx)
}
