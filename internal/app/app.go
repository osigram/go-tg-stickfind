package app

import (
	"github.com/NicoNex/echotron/v3"
	"go-tg-stickfind/internal/handlers"
	"go-tg-stickfind/internal/ocr"
	"go-tg-stickfind/internal/storage"
	"log/slog"
	"time"
)

type App struct {
	Logger      *slog.Logger
	HelpMessage string
	token       string
	Storage     storage.Storage
	OCRGetter   ocr.Getter
	echotron.API
}

func NewApp(
	logger *slog.Logger,
	storage storage.Storage,
	helpMessage string,
	token string,
	ocrGetter ocr.Getter,
) *App {
	return &App{
		Logger:      logger,
		Storage:     storage,
		HelpMessage: helpMessage,
		API:         echotron.NewAPI(token),
		OCRGetter:   ocrGetter,
		token:       token,
	}
}

func (app *App) newBot(chatId int64) echotron.Bot {
	return handlers.NewBot(chatId, app)
}

func (app *App) Start() {
	l := app.Logger.With("op", "app.Start")
	dsp := echotron.NewDispatcher(app.token, app.newBot)
	for {
		err := dsp.Poll()
		if err != nil {
			l.Error("Error to poll", slog.String("error", err.Error()))
		}
		// In case of connection issues wait 5 seconds before trying to reconnect.
		time.Sleep(5 * time.Second)
	}
}
