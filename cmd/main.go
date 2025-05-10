package main

import (
	"shop-go/internal/config"
	"shop-go/internal/pkg/logger"
	"shop-go/internal/server"
)

func main() {
	// Load configuration
	cfg := config.GetConfig()

	// Initialize logger
	logConfig := &logger.LogConfig{
		Level:      cfg.Logger.Level,
		Encoding:   cfg.Logger.Encoding,
		OutputPath: cfg.Logger.OutputPath,
	}
	logger.Init(logConfig)
	defer logger.Sync()

	// Create and start the server
	srv := server.NewServer(cfg)
	logger.Infof("Starting server on %s", cfg.Server.Port)
	if err := srv.Start(); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
