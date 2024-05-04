package main

import (
	"go-tg-stickfind/internal/app"
	"go-tg-stickfind/internal/config"
	"go-tg-stickfind/internal/handlers"
	"go-tg-stickfind/internal/log"
	"go-tg-stickfind/internal/ocr/google"
	"go-tg-stickfind/internal/storage/mongodb"
)

func main() {
	cfg := config.NewConfig()

	// Logger
	logger, writeCloser := log.NewLogger(cfg.BuildMode, cfg.IsConsoleLogger, cfg.LogFilePath)
	defer writeCloser.Close()
	logger.Info("Logger initialized")

	// Storage
	logger.Info("Initializing storage...")
	storage := mongodb.MustNewStorage(cfg.ConnectionString)

	// App
	logger.Info("Initializing app...")
	botApp := app.NewApp(logger, storage, cfg.HelloMessage, cfg.Token, google.NewOCR)

	botFactory := handlers.NewBotFactory(botApp)

	// Start bot
	logger.Info("Starting bot...")
	botApp.Start(botFactory)
}
