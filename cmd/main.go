package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/colinjuang/shop-go/internal/config"
	"github.com/colinjuang/shop-go/internal/pkg/logger"
	"github.com/colinjuang/shop-go/internal/server"
)

func main() {
	// 加载配置
	cfg := config.GetConfig()

	// 初始化 logger
	logConfig := &logger.LogConfig{
		Level:      cfg.Logger.Level,
		Encoding:   cfg.Logger.Encoding,
		OutputPath: cfg.Logger.OutputPath,
	}
	logger.Init(logConfig)
	defer logger.Sync()

	// 创建 server
	srv := server.NewServer(cfg)

	// 创建 server 错误通道
	serverErrors := make(chan error, 1)

	// 在 goroutine 中启动 server
	go func() {
		fmt.Println("Starting HTTP server...")
		serverErrors <- srv.Start()
	}()

	// 创建信号通道
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// 等待 server 错误或关闭信号
	select {
	case err := <-serverErrors:
		if err != nil {
			logger.Errorf("Server error: %v", err)
			fmt.Printf("======================================\n")
			fmt.Printf("Server error:\n%v\n", err)
			fmt.Printf("======================================\n")
			os.Exit(1)
		}
	case sig := <-shutdown:
		logger.Infof("Received shutdown signal: %v", sig)
		fmt.Printf("\nReceived shutdown signal: %v\n", sig)

		// 执行优雅关闭
		if err := srv.Shutdown(); err != nil {
			logger.Errorf("Graceful shutdown failed: %v", err)
			fmt.Printf("Graceful shutdown failed: %v\n", err)
			os.Exit(1)
		}

		logger.Info("Server shutdown successfully")
		fmt.Println("Server shutdown successfully")
	}
}
