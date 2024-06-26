package handlers

import (
	"github.com/NicoNex/echotron/v3"
	"go-tg-stickfind/internal/models"
	"log/slog"
)

func (b *Bot) Message(message *echotron.Message) {
	l := b.app.Logger.With(
		slog.String("op", "internal.handlers.Message"),
		slog.Int64("chatID", b.chatID),
		slog.String("message", message.Text),
	)

	err := b.app.Storage.RegisterUserIfNot(message.From.ID)
	if err != nil {
		l.Error("Error to register user", slog.String("error", err.Error()))
		return
	}

	cmd, err := models.NewInputCommand(message.Text)
	if err != nil {
		l.Debug("Error to parse command", slog.String("error", err.Error()))
		return
	}

	var answer string
	switch cmd.Command {
	case "start", "help":
		answer = b.Help(b.app.HelpMessage)
	case "feed":
		answer = b.FeedPackByName(message.From.ID, cmd.Params...)
	case "set_ocr_key":
		answer = b.SetOCRKey(message.From.ID, cmd.Params...)
	case "ocr_usage":
		answer = b.OCRUsage(message.From.ID)
	}

	if answer != "" {
		err := b.app.SendTextReply(answer, b.chatID, message.ID)
		if err != nil {
			l.Error("Error to send text reply", slog.String("error", err.Error()))
		}
	}
}
