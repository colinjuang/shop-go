package server

import (
	"shop-go/internal/config"
	"shop-go/internal/middleware"
	"shop-go/internal/pkg/logger"
	"shop-go/internal/pkg/minio"
	"shop-go/internal/pkg/redis"
	"shop-go/internal/repository"
	"shop-go/internal/routes"

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

// InitRoutes initializes the routes
func (s *Server) InitRoutes() {
	// Register all routes from the routes package
	routes.RegisterRoutes(s.router)
}

// Start starts the server
func (s *Server) Start() error {
	// Initialize database
	_, err := repository.InitDB(s.config)
	if err != nil {
		logger.Errorf("Failed to initialize database: %v", err)
		return err
	}
	logger.Info("Database connection established")

	// Initialize Redis
	_, err = redis.InitClient(&s.config.Redis)
	if err != nil {
		logger.Errorf("Failed to initialize Redis: %v", err)
		return err
	}
	logger.Info("Redis connection established")

	// Initialize MinIO
	_, err = minio.InitClient(&s.config.MinIO)
	if err != nil {
		logger.Errorf("Failed to initialize MinIO: %v", err)
		return err
	}
	logger.Info("MinIO connection established")

	// Initialize routes
	s.InitRoutes()
	logger.Info("Routes initialized")

	// Start server
	logger.Infof("Starting server on %s in %s mode", s.config.Server.Port, s.config.Server.Environment)
	return s.router.Run(s.config.Server.Port)
}
