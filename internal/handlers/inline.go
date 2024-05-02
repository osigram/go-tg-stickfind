package handlers

import (
	"github.com/NicoNex/echotron/v3"
	"log/slog"
)

const MaxStickers = 20 // limit of stickers in inline query by Telegram

func (b *Bot) ProcessInlineQuery(query *echotron.InlineQuery) {
	if len(query.Query) < 5 {
		return
	}

	l := b.app.Logger.With(
		slog.String("op", "internal.handler.ProcessInlineQuery"),
		slog.String("query", query.Query),
	)

	stickers, err := b.app.Storage.FindStickers(query.Query, MaxStickers)
	if err != nil {
		l.Error("Error to find stickers", slog.String("error", err.Error()))
		return
	}

	var results []echotron.InlineQueryResult
	for _, sticker := range stickers {
		results = append(results, echotron.InlineQueryResultCachedSticker{
			Type:          "sticker",
			ID:            sticker.ID,
			StickerFileID: sticker.FileID,
		})
	}

	_, err = b.app.AnswerInlineQuery(query.ID, results, nil)
	if err != nil {
		l.Error("Error to answer inline query", slog.String("error", err.Error()))
	}
}
