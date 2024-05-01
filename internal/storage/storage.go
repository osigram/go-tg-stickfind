package storage

import "go-tg-stickfind/internal/models"

type Storage interface {
	SetSticker(sticker models.Sticker) error
	FindSticker(text string) (*models.Sticker, error)
}
