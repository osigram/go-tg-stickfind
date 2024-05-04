package models

type User struct {
	ID            int64
	Key           string
	OCRUsageStats []OCRUsageStat
}

type OCRUsageStat struct {
	Usage int
	Month int
	Year  int
}
