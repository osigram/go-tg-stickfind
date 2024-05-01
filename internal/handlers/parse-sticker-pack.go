package handlers

import (
	"errors"
	"fmt"
	"github.com/NicoNex/echotron/v3"
	"go-tg-stickfind/internal/models"
	"go-tg-stickfind/internal/ocr"
)

type StickerParsingErr struct {
	StickerUniqueID string
	StickerSetName  string
	err             error
}

func (e *StickerParsingErr) Error() string {
	return fmt.Sprintf("error to parse sticker %v from pack %v: %v.", e.StickerUniqueID, e.StickerSetName, e.err)
}

func (b *Bot) parseStickerPack(pack *echotron.StickerSet) error {
	var err error

	for _, sticker := range pack.Stickers {
		err = b.parseSticker(&sticker)
		if err != nil {
			err = errors.Join(err)
		}
	}

	return err
}

func (b *Bot) parseSticker(sticker *echotron.Sticker) error {
	file, err := b.app.GetFile(sticker.FileID)
	if err != nil || !file.Ok {
		return &StickerParsingErr{
			StickerUniqueID: sticker.FileUniqueID,
			StickerSetName:  sticker.SetName,
			err:             fmt.Errorf("error to get file: %v", err),
		}
	}

	stickerBytes, err := b.app.DownloadFile(file.Result.FilePath)
	if err != nil {
		return &StickerParsingErr{
			StickerUniqueID: sticker.FileUniqueID,
			StickerSetName:  sticker.SetName,
			err:             fmt.Errorf("error to download file: %v", err),
		}
	}

	text, err := ocr.ParseImage(stickerBytes)
	if err != nil {
		return &StickerParsingErr{
			StickerUniqueID: sticker.FileUniqueID,
			StickerSetName:  sticker.SetName,
			err:             fmt.Errorf("error to parse image: %v", err),
		}
	}

	if err := b.app.Storage.SetSticker(models.Sticker{
		StickerID: sticker.FileID,
		Text:      text,
	}); err != nil {
		return &StickerParsingErr{
			StickerUniqueID: sticker.FileUniqueID,
			StickerSetName:  sticker.SetName,
			err:             fmt.Errorf("error to save sticker: %v", err),
		}
	}

	return nil
}
