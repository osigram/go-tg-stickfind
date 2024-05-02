package main

import (
	"go-tg-stickfind/internal/app"
	"go-tg-stickfind/internal/config"
	"go-tg-stickfind/internal/log"
	"go-tg-stickfind/internal/ocr/google"
)

func main() {
	cfg := config.NewConfig()

	// Logger
	logger, writeCloser := log.NewLogger(cfg.BuildMode, cfg.IsConsoleLogger, cfg.LogFilePath)
	defer writeCloser.Close()
	logger.Info("Logger initialized")

	// TODO: Storage
	logger.Info("Initializing storage...")

	// App
	logger.Info("Initializing app...")
	botApp := app.NewApp(logger, _, cfg.HelloMessage, cfg.Token, google.NewOCR)

	// Start bot
	logger.Info("Starting bot...")
	botApp.Start()
}
