package mongodb

import (
	"context"
	"go-tg-stickfind/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (s *Storage) SetUserKey(userID int64, key string) error {
	db := s.db

	filter := bson.D{{"_id", userID}}
	update := bson.D{{"$set", bson.D{{"key", key}}}}
	_, err := db.Collection("users").
		UpdateOne(
			context.Background(),
			filter,
			update,
		)

	return err
}

func (s *Storage) ProcessOCRUsage(userID int64, OCRUsage int) error {
	db := s.db
	year, month, _ := time.Now().Date()

	filter := bson.D{
		{"_id", userID},
		{"ocr_usage_stat.month", int(month)},
		{"ocr_usage_stat.year", year},
	}
	update := bson.D{{"$inc", bson.D{{"ocr_usage_stat.$.usage", OCRUsage}}}}
	result, err := db.Collection("users").
		UpdateOne(
			context.Background(),
			filter,
			update,
		)
	if err != nil {
		return err
	}

	if result.ModifiedCount != 0 {
		return nil
	}

	// If no document was updated, we need to add a new month to the array
	filter = bson.D{{"_id", userID}}
	update = bson.D{{
		"$addToSet",
		bson.D{
			{"ocr_usage_stat",
				models.OCRUsageStat{Usage: OCRUsage, Month: int(month), Year: year},
			},
		},
	}}
	_, err = db.Collection("users").
		UpdateOne(
			context.Background(),
			filter,
			update,
		)

	return err
}

func (s *Storage) GetUser(userID int64) (*models.User, error) {
	db := s.db

	var user models.User
	err := db.Collection("users").
		FindOne(context.Background(), bson.M{"_id": userID}).
		Decode(&user)

	return &user, err
}

func (s *Storage) RegisterUserIfNot(userID int64) error {
	db := s.db

	year, month, _ := time.Now().Date()
	user := models.User{
		Key: "",
		OCRUsageStats: []models.OCRUsageStat{
			{Usage: 0, Month: int(month), Year: year},
		},
	}

	_, err := db.Collection("users").
		UpdateByID(
			context.Background(),
			userID, bson.M{"$setOnInsert": user},
			options.Update().SetUpsert(true),
		)

	return err
}
