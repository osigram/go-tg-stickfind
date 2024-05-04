package mongodb

import (
	"context"
	"go-tg-stickfind/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Storage) GetSticker(id string) (*models.Sticker, error) {
	db := s.db

	var sticker models.Sticker
	err := db.Collection("stickers").
		FindOne(context.Background(), bson.D{{"_id", id}}).
		Decode(&sticker)

	return &sticker, err
}

func (s *Storage) SetSticker(sticker models.Sticker) error {
	db := s.db

	_, err := db.Collection("stickers").
		InsertOne(context.Background(), sticker)

	return err
}

func (s *Storage) FindStickersByText(text string, maxNum int) ([]models.Sticker, error) {
	db := s.db

	cursor, err := db.Collection("stickers").
		Find(
			context.Background(),
			bson.D{{"$text", bson.D{{"$search", text}}}},
			options.Find().SetLimit(int64(maxNum)),
		)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var stickers []models.Sticker
	err = cursor.All(context.Background(), &stickers)

	return stickers, err
}
