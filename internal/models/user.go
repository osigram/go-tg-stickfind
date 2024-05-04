package models

type User struct {
	ID            int64          `bson:"_id,omitempty"`
	Key           string         `bson:"key"`
	OCRUsageStats []OCRUsageStat `bson:"ocr_usage_stat"`
}

type OCRUsageStat struct {
	Usage int `bson:"usage"`
	Month int `bson:"month"`
	Year  int `bson:"year"`
}
