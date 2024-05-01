package storagemock

import (
	"errors"
	"go-tg-stickfind/internal/models"
	"strings"
	"sync"
)

type StorageMock struct {
	data map[string]string
	mx   *sync.RWMutex
}

func NewStorageMock() *StorageMock {
	return &StorageMock{
		data: make(map[string]string),
		mx:   &sync.RWMutex{},
	}
}

func (s *StorageMock) SetSticker(sticker models.Sticker) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.data[sticker.StickerID] = sticker.Text
	return nil
}

func (s *StorageMock) FindSticker(text string) (*models.Sticker, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	for id, t := range s.data {
		if strings.Contains(t, text) {
			return &models.Sticker{
				StickerID: id,
				Text:      t,
			}, nil
		}
	}

	return nil, errors.New("sticker not found")
}
