package bot

import (
	"context"
	"learning_bot/core"
	"learning_bot/handlers"
	"learning_bot/misc"

	"github.com/go-telegram/bot"
)

func SetupBot(ctx context.Context) *bot.Bot {
	b, err := bot.New(
		core.Cfg.BotConfig.Token,
		bot.WithDefaultHandler(handlers.DefaultHandler),
		bot.WithMiddlewares(CollectUserIfNotExists),
	)
	misc.Must(err)

	wh_info, err := b.GetWebhookInfo(ctx)
	misc.Must(err)

	if wh_info.URL != "" {
		b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{DropPendingUpdates: true})
	}
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.StartCommandHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, handlers.HelpCommandHandler)

	return b
}
