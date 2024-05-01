package main

import (
	"go-tg-stickfind/internal/bot"
	"go-tg-stickfind/internal/config"
	"go-tg-stickfind/internal/log"
)

func main() {
	cfg := config.NewConfig()

	// Logger
	logger, writeCloser := log.NewLogger(cfg.BuildMode, cfg.IsConsoleLogger, cfg.LogFilePath)
	defer writeCloser.Close()
	logger.Info("Logger initialized")

	logger.Info("Starting bot...")
	botApp := bot.NewApp(logger, cfg.HelloMessage, cfg.Token)
	botApp.Start()
}
