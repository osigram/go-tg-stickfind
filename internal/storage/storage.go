package storage

import "go-tg-stickfind/internal/models"

type Storage interface {
	FindSticker(ID string) (*models.Sticker, error)
	SetSticker(sticker models.Sticker) error
	FindStickers(text string, maxNum int) ([]models.Sticker, error)
	SetUserKey(userID int64, key string) error
	GetUserKey(userID int64) (string, error)
	RegisterUser(userID int64) error
}
