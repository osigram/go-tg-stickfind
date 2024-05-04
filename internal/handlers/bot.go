package handlers

import (
	"github.com/NicoNex/echotron/v3"
	"go-tg-stickfind/internal/app"
)

type Bot struct {
	chatID int64
	app    *app.App
}

func NewBot(chatID int64, app *app.App) *Bot {
	return &Bot{
		chatID: chatID,
		app:    app,
	}
}

func NewBotFactory(app *app.App) echotron.NewBotFn {
	return func(chatID int64) echotron.Bot {
		return NewBot(chatID, app)
	}
}
