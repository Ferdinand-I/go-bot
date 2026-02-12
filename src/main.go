package main

import (
	"context"
	"learning_bot/bot"
	"learning_bot/core"
	"learning_bot/misc"
	"learning_bot/storage"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	db := storage.ConnectDB(core.Cfg.DBConfig)
	defer db.Close()

	ctx := misc.NewContext(context.Background(), db)
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	b := bot.SetupBot(ctx)

	log.Println("bot started")
	b.Start(ctx)
}
