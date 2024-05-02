package handlers

import (
	"github.com/NicoNex/echotron/v3"
	"go-tg-stickfind/internal/models"
	"log/slog"
)

func (b *Bot) Update(update *echotron.Update) {
	l := b.app.Logger.With(
		slog.String("op", "internal.handler.Update"),
		slog.Int64("chatID", b.chatID),
	)

	cmd, err := models.NewInputCommand(update.Message.Text)
	if err != nil {
		l.Debug("Error to parse command", slog.String("error", err.Error()))
		return
	}

	ocr, err := b.app.GetOCR(update.Message.From.ID)
	if err != nil {
		l.Error("Error to get ocr", slog.String("error", err.Error()))
		return
	}

	var answer string
	switch cmd.Command {
	case "start", "help":
		answer = b.Help(b.app.HelpMessage)
	case "feed":
		answer = b.FeedPackByName(ocr, cmd.Params...)
	}

	if answer != "" {
		err := b.app.SendTextReply(answer, b.chatID, update.Message.ID)
		if err != nil {
			l.Error("Error to send text reply", slog.String("error", err.Error()))
		}
	}
}
