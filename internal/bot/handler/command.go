package handler

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) StartCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Hello, now you are in my DB! 😈",
	})
	if err != nil {
		log.Printf("failed to send start message: %v", err)
	}
}

func (h *Handler) HelpCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Help message",
	})
	if err != nil {
		log.Printf("failed to send help message: %v", err)
	}
}
