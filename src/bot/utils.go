package bot

import (
	"context"
	"learning_bot/core"
	"learning_bot/handlers"
	"log"

	"github.com/go-telegram/bot"
)

func SetupBot(ctx context.Context) *bot.Bot {
	b, err := bot.New(core.Cfg.BotConfig.Token, bot.WithDefaultHandler(handlers.DefaultHandler))
	if err != nil {
		log.Fatal(err)
	}

	wh_info, err := b.GetWebhookInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if wh_info.URL != "" {
		b.DeleteWebhook(ctx, &bot.DeleteWebhookParams{DropPendingUpdates: true})
	}
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.StartCommandHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, handlers.HelpCommandHandler)

	return b
}
