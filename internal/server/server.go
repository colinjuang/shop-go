package server

import (
	"fmt"

	"github.com/colinjuang/shop-go/internal/config"
	"github.com/colinjuang/shop-go/internal/middleware"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"github.com/colinjuang/shop-go/internal/pkg/minio"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/colinjuang/shop-go/internal/router"
	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server
type Server struct {
	config *config.Config
	router *gin.Engine
}

// NewServer creates a new server
func NewServer(cfg *config.Config) *Server {
	// Set Gin mode based on environment
	if cfg.Server.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a custom gin.Engine with our logger
	router := gin.New()

	// Use our own logger and recovery middleware
	router.Use(middleware.ZapLogger())
	router.Use(gin.Recovery())

	return &Server{
		config: cfg,
		router: router,
	}
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

	// Initialize Redis
	_, err = redis.InitClient(&s.config.Redis)
	if err != nil {
		fmt.Printf("Failed to initialize Redis: %v\n", err)
		return err
	}
	fmt.Println("Redis connection established")

	// Initialize MinIO
	_, err = minio.InitClient(&s.config.MinIO)
	if err != nil {
		fmt.Printf("Failed to initialize MinIO: %v\n", err)
		return err
	}
	fmt.Println("MinIO connection established")

	// Initialize routes
	router.RegisterRouter(s.router)
	fmt.Println("Routes initialized")

	// Start server
	fmt.Printf("Starting server on %s in %s mode\n", s.config.Server.Port, s.config.Server.Environment)
	return s.router.Run(s.config.Server.Port)
}
