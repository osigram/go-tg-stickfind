package main

import (
	"go-tg-stickfind/internal/app"
	"go-tg-stickfind/internal/config"
	"go-tg-stickfind/internal/log"
	"go-tg-stickfind/internal/storage/storagemock"
)

func main() {
	cfg := config.NewConfig()

	// Logger
	logger, writeCloser := log.NewLogger(cfg.BuildMode, cfg.IsConsoleLogger, cfg.LogFilePath)
	defer writeCloser.Close()
	logger.Info("Logger initialized")

	// Storage
	logger.Info("Initializing storage...")
	storage := storagemock.NewStorageMock()

	// App
	logger.Info("Initializing app...")
	botApp := app.NewApp(logger, storage, cfg.HelloMessage, cfg.Token)

	// Start bot
	logger.Info("Starting bot...")
	botApp.Start()
}
