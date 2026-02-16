package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	tgbot "learning_bot/internal/bot"
	"learning_bot/internal/config"
	"learning_bot/internal/storage"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := storage.ConnectDB(cfg.DB.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	userRepo := storage.NewUserRepo(db)

	b, err := tgbot.New(cfg.Bot.Token, userRepo)
	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := tgbot.CleanupWebhook(ctx, b); err != nil {
		log.Fatalf("failed to cleanup webhook: %v", err)
	}

	log.Println("bot started")
	b.Start(ctx)
}
