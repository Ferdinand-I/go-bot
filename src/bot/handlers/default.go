package handlers

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.From.IsBot {
		return
	}

	if update.Message == nil {
		return
	}

	_, err := b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: update.Message.ID,
			Text:   "Send `/help` to get all available commands",
		},
	)
	if err != nil {
		log.Printf("failed to send default message: %v", err)
	}
}
