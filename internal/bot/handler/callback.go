package handler

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func callbackQueryPrefix(s string) string {
	return strings.SplitN(s, ":", 2)[0]
}

func ContainsPhoto(update *models.Update) bool {
	return update.Message != nil && update.Message.Photo != nil
}

func IsDeleteMessageCQ(update *models.Update) bool {
	return update.CallbackQuery != nil &&
		callbackQueryPrefix(update.CallbackQuery.Data) == "deleteMessage"
}

func IsPinMessageCQ(update *models.Update) bool {
	return update.CallbackQuery != nil &&
		callbackQueryPrefix(update.CallbackQuery.Data) == "pinMessage"
}

func (h *Handler) DeleteMessageCQ(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := strings.Split(update.CallbackQuery.Data, ":")
	chatID, messageID := data[1], data[2]

	messageIDInt, err := strconv.Atoi(messageID)
	if err != nil {
		log.Printf("failed to parse message ID: %v", err)
		return
	}

	deleted, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    chatID,
		MessageID: messageIDInt,
	})
	if err != nil || !deleted {
		log.Printf("failed to delete message: %v", err)
	}
}

func (h *Handler) PinMessageCQ(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := strings.Split(update.CallbackQuery.Data, ":")
	chatID, messageID := data[1], data[2]

	messageIDInt, err := strconv.Atoi(messageID)
	if err != nil {
		log.Printf("failed to parse message ID: %v", err)
		return
	}

	_, err = b.PinChatMessage(ctx, &bot.PinChatMessageParams{
		ChatID:    chatID,
		MessageID: messageIDInt,
	})
	if err != nil {
		log.Printf("failed to pin message: %v", err)
	}
}
