package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/colinjuang/shop-go/internal/app/middleware"
	"github.com/colinjuang/shop-go/internal/app/router"
	"github.com/colinjuang/shop-go/internal/config"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"github.com/colinjuang/shop-go/internal/pkg/minio"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server
type Server struct {
	config     *config.Config
	router     *gin.Engine
	httpServer *http.Server
}

// NewServer creates a new server
func NewServer(cfg *config.Config) *Server {
	// 设置 gin 模式
	if cfg.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建一个带有我们日志记录器的自定义 gin.Engine
	router := gin.New()

	// 使用我们自己的日志记录器和恢复中间件
	router.Use(middleware.ZapLogger())
	router.Use(gin.Recovery())

	return &Server{
		config: cfg,
		router: router,
	}
}

// GetRouter returns the router instance for external access if needed
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}

// GetConfig returns the server configuration
func (s *Server) GetConfig() *config.Config {
	return s.config
}

// Start starts the server
func (s *Server) Start() error {
	// Initialize database
	_, err := database.InitDB(s.config)
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		return err
	}
	fmt.Println("Database connection established")

	// Add database migration
	// if err := database.AutoMigrate(); err != nil {
	// 	return err
	// }

	// 初始化 Redis
	_, err = redis.InitClient(&s.config.Redis)
	if err != nil {
		fmt.Printf("Failed to initialize Redis: %v\n", err)
		return err
	}
	fmt.Println("Redis connection established")

	// 初始化 MinIO
	_, err = minio.InitClient(&s.config.MinIO)
	if err != nil {
		fmt.Printf("Failed to initialize MinIO: %v\n", err)
		return err
	}
	fmt.Println("MinIO connection established")

	// 初始化路由
	router.RegisterRouter(s.router)
	fmt.Println("Routes initialized")

	// 创建 HTTP 服务器
	s.httpServer = &http.Server{
		Addr:           s.config.Server.Port,
		Handler:        s.router,
		MaxHeaderBytes: 1 << 20, // 1MB
		// 设置超时
		ReadTimeout:  time.Duration(s.config.Server.ReadTimeout) * time.Second,  // 读取超时
		WriteTimeout: time.Duration(s.config.Server.WriteTimeout) * time.Second, // 写入超时
		IdleTimeout:  time.Duration(s.config.Server.IdleTimeout) * time.Second,  // 空闲超时
	}

	// 启动服务器
	fmt.Printf("Starting server on %s in %s mode\n", s.config.Server.Port, s.config.Server.Environment)
	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown() error {
	fmt.Println("Starting graceful shutdown...")

	// 创建带有超时上下文的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 优雅关闭 HTTP 服务器
	if s.httpServer != nil {
		fmt.Println("Shutting down HTTP server...")
		if err := s.httpServer.Shutdown(ctx); err != nil {
			fmt.Printf("HTTP server shutdown error: %v\n", err)
			return fmt.Errorf("failed to shutdown HTTP server: %w", err)
		}
		fmt.Println("HTTP server shutdown complete")
	}

	// 关闭数据库连接
	if db := database.GetDB(); db != nil {
		fmt.Println("Closing database connections...")
		if sqlDB, err := db.DB(); err == nil {
			if err := sqlDB.Close(); err != nil {
				fmt.Printf("Database close error: %v\n", err)
				return fmt.Errorf("failed to close database: %w", err)
			}
		}
		fmt.Println("Database connections closed")
	}

	// 关闭 Redis 连接
	if redisClient := redis.GetClient(); redisClient != nil {
		fmt.Println("Closing Redis connections...")
		if err := redisClient.Close(); err != nil {
			fmt.Printf("Redis close error: %v\n", err)
			return fmt.Errorf("failed to close Redis: %w", err)
		}
		fmt.Println("Redis connections closed")
	}

	// 注意：MinIO 客户端不需要显式关闭，因为它使用 HTTP 连接
	// 这些连接由 HTTP 传输处理

	fmt.Println("Graceful shutdown completed successfully")
	return nil
}
