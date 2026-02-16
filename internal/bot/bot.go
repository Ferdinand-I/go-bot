package bot

import (
	"context"
	"fmt"

	"learning_bot/internal/bot/handler"
	"learning_bot/internal/storage"

	"github.com/go-telegram/bot"
)

func New(token string, userRepo *storage.UserRepo) (*bot.Bot, error) {
	h := handler.New(userRepo)

	b, err := bot.New(
		token,
		bot.WithDefaultHandler(h.Default),
		bot.WithMiddlewares(CollectUserIfNotExists(userRepo)),
	)
	if err != nil {
		return nil, fmt.Errorf("creating bot: %w", err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, h.StartCommand)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, h.HelpCommand)
	b.RegisterHandlerMatchFunc(handler.ContainsPhoto, h.PhotoMessage)
	b.RegisterHandlerMatchFunc(handler.IsDeleteMessageCQ, h.DeleteMessageCQ)
	b.RegisterHandlerMatchFunc(handler.IsPinMessageCQ, h.PinMessageCQ)

	return b, nil
}

func CleanupWebhook(ctx context.Context, b *bot.Bot) error {
	whInfo, err := b.GetWebhookInfo(ctx)
	if err != nil {
		return fmt.Errorf("getting webhook info: %w", err)
	}

	if whInfo.URL != "" {
		_, _ = b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{DropPendingUpdates: true})
	}

	return nil
}
