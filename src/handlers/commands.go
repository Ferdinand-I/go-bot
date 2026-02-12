package handlers

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartCommandHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Hello, now you are in my DB! 😈",
		},
	)

	if err != nil {
		log.Panic(err)
	}
}

func HelpCommandHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Help message"},
	)

	if err != nil {
		log.Panic(err)
	}
}
