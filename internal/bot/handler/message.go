package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) PhotoMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         "Удали сейчас же!",
					CallbackData: fmt.Sprintf("deleteMessage:%d:%d", update.Message.Chat.ID, update.Message.ID),
				},
			},
			{
				{
					Text:         "Кинь в закреп!",
					CallbackData: fmt.Sprintf("pinMessage:%d:%d", update.Message.Chat.ID, update.Message.ID),
				},
			},
		},
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "I recieved a photo from you! What do you want me to do with it?",
		ReplyMarkup: kb,
	})
	if err != nil {
		log.Printf("failed to send photo message: %v", err)
	}
}
