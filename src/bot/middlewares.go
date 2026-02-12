package bot

import (
	"context"
	"learning_bot/misc"
	"learning_bot/storage/db_models"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/jmoiron/sqlx"
)

// CollectUserIfNotExists middleware checks if the user exists in the database and creates a new user if not.
func CollectUserIfNotExists(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		db := misc.DBFromCtx(ctx)

		go func(db *sqlx.DB, update *models.Update) {
			user := db_models.GetUserByTgID(db, update.Message.From.ID)

			if user == nil {
				err := db_models.CreateUser(
					db,
					&db_models.User{
						TGID:      update.Message.From.ID,
						FirstName: update.Message.From.FirstName,
						Username:  update.Message.From.Username,
					},
				)
				misc.Must(err)
			}
		}(db, update)

		next(ctx, b, update)
	}
}
