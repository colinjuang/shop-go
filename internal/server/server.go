package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/colinjuang/shop-go/internal/config"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"github.com/colinjuang/shop-go/internal/pkg/minio"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	srvCtx *ServerContext
)

// Server represents the HTTP server
type ServerContext struct {
	config     *config.Config
	httpServer *http.Server
	DB         *gorm.DB
	Redis      *redis.Client
	Minio      *minio.Client
}

func GetServer() *ServerContext {
	return srvCtx
}

// NewServer creates a new server
func NewServerContext(cfg *config.Config) *ServerContext {
	// 初始化数据库
	// fmt.Printf("cfg.DatabaseConf: %+v\n", cfg.DatabaseConf)
	db, err := database.InitDB(&cfg.DatabaseConf)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
		return nil
	}
	fmt.Println("Database connection established")

	// 初始化 Redis
	redisClient, err := redis.InitClient(&cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v\n", err)
		return nil
	}
	fmt.Println("Redis connection established")

	// 初始化 MinIO
	minioClient, err := minio.InitClient(&cfg.MinIO)
	if err != nil {
		log.Fatalf("Failed to initialize MinIO: %v\n", err)
		return nil
	}
	fmt.Println("MinIO connection established")

	srvCtx = &ServerContext{
		config: cfg,
		DB:     db,
		Redis:  redisClient,
		Minio:  minioClient,
	}
	return srvCtx
}

// GetConfig returns the server configuration
func (s *ServerContext) GetConfig() *config.Config {
	return s.config
}

// Start starts the server
func (s *ServerContext) Start(router *gin.Engine) error {
	// 创建 HTTP 服务器
	s.httpServer = &http.Server{
		Addr:           s.config.Server.Port,
		Handler:        router,
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
func (s *ServerContext) Shutdown() error {
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
	fmt.Println("Closing database connections...")
	if err := database.Close(s.DB); err != nil {
		fmt.Printf("Database close error: %v\n", err)
		return fmt.Errorf("failed to close database: %w", err)
	}
	fmt.Println("Database connections closed")

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
