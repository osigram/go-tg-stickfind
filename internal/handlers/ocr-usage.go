package handlers

import (
	"go-tg-stickfind/internal/models"
	"log/slog"
	"strconv"
	"time"
)

func (b *Bot) OCRUsage(userID int64) string {
	l := b.app.Logger.With(slog.String("op", "internal.handlers.OCRUsage"))

	user, err := b.app.Storage.GetUser(userID)
	if err != nil {
		l.Error("Error to get user", slog.String("error", err.Error()))
		return "Error to get user."
	}

	var OCRUsage models.OCRUsageStat
	year, month, _ := time.Now().Date()
	for _, stat := range user.OCRUsageStats {
		if stat.Month == int(month) && stat.Year == year {
			OCRUsage = stat
		}
	}

	answer := "Your OCR usage for this month is " + strconv.Itoa(OCRUsage.Usage) + "."

	return answer
}
