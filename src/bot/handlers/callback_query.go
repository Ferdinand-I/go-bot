package handlers

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DeleteMessageCQHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := strings.Split(update.CallbackQuery.Data, ":")
	chatId, messageId := data[1], data[2]

	messageIdInt, err := strconv.Atoi(messageId)
	if err != nil {
		log.Print("Error while converting message ID to integer")
		return
	}

	deleted, error := b.DeleteMessage(
		ctx,
		&bot.DeleteMessageParams{ChatID: chatId, MessageID: messageIdInt},
	)

	if error != nil || deleted == false {
		log.Print("Error while deleting message")
	}
}

func PinMessageCQHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := strings.Split(update.CallbackQuery.Data, ":")
	chatId, messageId := data[1], data[2]

	messageIdInt, err := strconv.Atoi(messageId)
	if err != nil {
		log.Print("Error while converting message ID to integer")
		return
	}

	_, error := b.PinChatMessage(
		ctx,
		&bot.PinChatMessageParams{ChatID: chatId, MessageID: messageIdInt},
	)
	if error != nil {
		log.Print("Error while pinning message")
	}
}
