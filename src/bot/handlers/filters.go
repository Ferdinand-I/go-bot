package handlers

import (
	"learning_bot/misc"

	"github.com/go-telegram/bot/models"
)

func ContainsPhoto(update *models.Update) bool {
	if update.Message != nil && update.Message.Photo != nil {
		return true
	}

	return false
}

func DeleteMessageCQ(update *models.Update) bool {
	if update.CallbackQuery != nil &&
		misc.GetCallbackQueryPrefix(update.CallbackQuery.Data) == "deleteMessage" {
		return true
	}

	return false
}

func PinMessageCQ(update *models.Update) bool {
	if update.CallbackQuery != nil &&
		misc.GetCallbackQueryPrefix(update.CallbackQuery.Data) == "pinMessage" {
		return true
	}

	return false
}
