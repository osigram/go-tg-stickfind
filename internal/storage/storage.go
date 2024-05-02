package storage

import "go-tg-stickfind/internal/models"

type Storage interface {
	SetSticker(sticker models.Sticker) error
	FindSticker(text string) (*models.Sticker, error)
	SetUserKey(userID int64, key string) error
	GetUserKey(userID int64) (string, error)
	RegisterUser(userID int64) error
}
