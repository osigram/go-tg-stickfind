package handlers

import (
	"fmt"
	"go-tg-stickfind/internal/ocr"
	"log/slog"
	"strings"
)

func (b *Bot) FeedPackByName(ocr ocr.OCR, packs ...string) string {
	l := b.app.Logger.With(slog.String("op", "internal.handlers.FeedPackByName"))
	var answers []string

	for _, stickerPackName := range packs {
		pack, err := b.app.GetStickerSet(stickerPackName)
		if err != nil {
			l.Debug("Error to get sticker pack", slog.String("error", err.Error()), slog.String("stickerPackName", stickerPackName))
			answers = append(answers, fmt.Sprintf("Error to get sticker pack: %v.", stickerPackName))

			continue
		}

		if !pack.Ok || pack.Result.IsAnimated || pack.Result.IsVideo || pack.Result.StickerType != "regular" {
			l.Debug("Sticker pack is not static", slog.String("stickerPackName", stickerPackName))
			answers = append(answers, fmt.Sprintf("Sticker pack %v is not regular.", stickerPackName))

			continue
		}

		if err := b.parseStickerPack(pack.Result, ocr); err != nil {
			answers = append(answers, fmt.Sprintf("Error while parsing sticker pack: %v.", stickerPackName))
			l.Error("Error while parsing sticker pack", slog.String("error", err.Error()), slog.String("stickerPackName", stickerPackName))
		}
	}

	return strings.Join(answers, "\n")
}
