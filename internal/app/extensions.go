package app

import (
	"fmt"
	"github.com/NicoNex/echotron/v3"
	"go-tg-stickfind/internal/ocr"
)

func (app *App) SendTextReply(text string, chatID int64, replyID int) error {
	_, err := app.SendMessage(text, chatID, &echotron.MessageOptions{
		ReplyParameters: echotron.ReplyParameters{MessageID: replyID},
	})

	return err
}

func (app *App) GetOCR(userID int64) (ocr.OCR, error) {
	key, err := app.Storage.GetUserKey(userID)
	if err != nil {
		return nil, fmt.Errorf("error to get user ocr key: %v", err)
	}

	ocrObject, err := app.OCRGetter(key)
	if err != nil {
		return nil, fmt.Errorf("error to get ocr: %v", err)
	}

	return ocrObject, err
}
