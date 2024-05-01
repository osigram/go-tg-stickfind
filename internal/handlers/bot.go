package handlers

import "go-tg-stickfind/internal/app"

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
