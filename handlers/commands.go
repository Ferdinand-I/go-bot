package handlers

import (
	"context"
	"learning_bot/misc"
	"learning_bot/storage/db_models"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartCommandHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	db := misc.DBFromCtx(ctx)

	user := db_models.GetUserByTgID(db, update.Message.From.ID)

	if user == nil {
		db_models.CreateUser(
			db,
			&db_models.User{
				TGID:      update.Message.From.ID,
				FirstName: update.Message.From.FirstName,
				Username:  update.Message.From.Username,
			},
		)
	}

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
