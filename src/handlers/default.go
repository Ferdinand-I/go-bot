package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// DefaultHandler is the default handler for the bot.
// Used to filter messages that are not handled by other handlers.

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil && update.Message.Photo != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "You sent a photo!",
		})
	}
}
