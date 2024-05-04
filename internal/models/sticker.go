package models

type Sticker struct {
	ID     string `bson:"_id"`
	FileID string `bson:"file_id"`
	Text   string `bson:"text"`
}
