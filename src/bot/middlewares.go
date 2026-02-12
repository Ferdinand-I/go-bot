package bot

import (
	"context"
	"learning_bot/misc"
	"learning_bot/storage/dbmodels"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// CollectUserIfNotExists middleware checks if the user exists in the database and creates a new user if not.
func CollectUserIfNotExists(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil || update.Message.From == nil {
			next(ctx, b, update)
			return
		}

		db := misc.DBFromCtx(ctx)
		from := update.Message.From
		userRepo := &dbmodels.UserRepo{DB: db}

		exists, err := userRepo.TgUserExists(from.ID)
		if err != nil {
			log.Printf("failed to check user existing: %v", err)
			next(ctx, b, update)
			return
		}

		if !exists {
			err := userRepo.Create(&dbmodels.User{
				TGID:      from.ID,
				FirstName: from.FirstName,
				Username:  from.Username,
			})
			if err != nil {
				log.Printf("failed to create user: %v", err)
			}
		}

		next(ctx, b, update)
	}
}
