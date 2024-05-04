package storage

import "go-tg-stickfind/internal/models"

type Storage interface {
	GetSticker(ID string) (*models.Sticker, error)
	SetSticker(sticker models.Sticker) error
	FindStickersByText(text string, maxNum int) ([]models.Sticker, error)
	SetUserKey(userID int64, key string) error
	ProcessOCRUsage(userID int64, OCRUsage int) error
	GetUser(userID int64) (*models.User, error)
	RegisterUser(userID int64) error
}
