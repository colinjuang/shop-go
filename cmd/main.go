package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/colinjuang/shop-go/internal/app/router"
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
	srv := server.NewServerContext(cfg)
	// 创建 router
	r := router.NewRouter(cfg)

	// 创建 server 错误通道
	serverErrors := make(chan error, 1)

	// 在 goroutine 中启动 server
	go func() {
		fmt.Println("Starting HTTP server...")
		router.RegisterRouter(r)
		fmt.Println("Routes initialized")
		serverErrors <- srv.Start(r)
	}()

	// 创建信号通道
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// 等待 server 错误或关闭信号
	select {
	case err := <-serverErrors:
		if err != nil {
			fmt.Println("======================================")
			fmt.Printf("Server error:\n%v\n", err)
			fmt.Println("======================================")
			os.Exit(1)
		}
	case sig := <-shutdown:
		fmt.Printf("Received shutdown signal: %v\n", sig)

		// 执行优雅关闭
		if err := srv.Shutdown(); err != nil {
			fmt.Printf("Graceful shutdown failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Server shutdown successfully")
	}
}
