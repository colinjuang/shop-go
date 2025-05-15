package main

import (
	"fmt"
	"os"
	"github.com/colinjuang/shop-go/internal/config"
	"github.com/colinjuang/shop-go/internal/pkg/logger"
	"github.com/colinjuang/shop-go/internal/server"
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
	fmt.Printf("Starting server on %s\n", cfg.Server.Port)
	if err := srv.Start(); err != nil {
		fmt.Println("======================================")
		fmt.Printf("Failed to start server:\n%v\n", err)
		fmt.Println("======================================")
		os.Exit(1)
	}
}
