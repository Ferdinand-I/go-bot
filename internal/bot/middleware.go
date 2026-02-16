package bot

import (
	"context"
	"log"

	"learning_bot/internal/storage"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func CollectUserIfNotExists(userRepo *storage.UserRepo) bot.Middleware {
	return func(next bot.HandlerFunc) bot.HandlerFunc {
		return func(ctx context.Context, b *bot.Bot, update *models.Update) {
			if update.Message == nil || update.Message.From == nil {
				next(ctx, b, update)
				return
			}

			from := update.Message.From

			exists, err := userRepo.TgUserExists(from.ID)
			if err != nil {
				log.Printf("failed to check user existence: %v", err)
				next(ctx, b, update)
				return
			}

			if !exists {
				err := userRepo.Create(&storage.User{
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
}
