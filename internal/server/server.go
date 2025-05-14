package server

import (
	"shop-go/internal/config"
	"shop-go/internal/handler"
	"shop-go/internal/middleware"
	"shop-go/internal/pkg/logger"
	"shop-go/internal/pkg/minio"
	"shop-go/internal/pkg/redis"
	"shop-go/internal/repository"
	"time"

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
	// Create handlers
	userHandler := handler.NewUserHandler()
	addressHandler := handler.NewAddressHandler()
	bannerHandler := handler.NewBannerHandler()
	categoryHandler := handler.NewCategoryHandler()
	promotionHandler := handler.NewPromotionHandler()
	productHandler := handler.NewProductHandler()
	cartHandler := handler.NewCartHandler()
	orderHandler := handler.NewOrderHandler()
	uploadHandler := handler.NewUploadHandler()
	reportHandler := handler.NewReportHandler()

	// Set up CORS
	s.router.Use(middleware.CORSMiddleware())

	// Set up global rate limiting - 100 requests per minute
	s.router.Use(middleware.RateLimitMiddleware(100, 1*time.Minute))

	// API group
	api := s.router.Group("/api")

	// Public endpoints (no auth required)
	// Home page
	api.GET("/home/banners", bannerHandler.GetBanners)
	api.GET("/home/category/level1", categoryHandler.GetCategories)
	api.GET("/home/promotions", promotionHandler.GetPromotions)
	api.GET("/home/recommend", productHandler.GetRecommendProducts)
	api.GET("/home/hot", productHandler.GetHotProducts)

	// Categories
	api.GET("/categories", categoryHandler.GetAllCategories)
	api.GET("/categories/:id/subs", categoryHandler.GetSubCategories)
	api.GET("/categories/tree", categoryHandler.GetCategoryTree)

	// Products
	api.GET("/products", productHandler.GetProducts)
	api.GET("/products/:id", productHandler.GetProductDetail)

	// Reports (public)
	api.GET("/reports/catalog", reportHandler.GetProductCatalog)
	api.GET("/reports/export", reportHandler.ExportProducts)

	// User login
	// api.GET("/user/login/:code", userHandler.GetWechatLoginCode)
	api.GET("/user/login", userHandler.WechatLogin)

	// Protected endpoints (auth required)
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())

	// User
	auth.GET("/users/info", userHandler.GetUserInfo)
	auth.POST("/users/update", userHandler.UpdateUserInfo)

	// Address
	auth.POST("/address/add", addressHandler.CreateAddress)
	auth.GET("/address/list", addressHandler.GetAddressList)
	auth.GET("/address/:id", addressHandler.GetAddressDetail)
	auth.POST("/address/:id/update", addressHandler.UpdateAddress)
	auth.GET("/address/:id/delete", addressHandler.DeleteAddress)

	// Cart
	auth.GET("/cart/add", cartHandler.AddToCart)
	auth.GET("/cart/list", cartHandler.GetCartList)
	auth.GET("/cart/update", cartHandler.UpdateCartItemStatus)
	auth.GET("/cart/update-all", cartHandler.UpdateAllCartItemStatus)
	auth.GET("/cart/delete", cartHandler.DeleteCartItem)

	// Order
	auth.GET("/order/detail", orderHandler.GetOrderDetail)
	auth.GET("/order/address", orderHandler.GetOrderAddress)
	auth.POST("/order/submit", orderHandler.CreateOrder)
	auth.GET("/order/pay", orderHandler.GetWechatPayInfo)
	auth.GET("/order/pay/status", orderHandler.CheckWechatPayStatus)
	auth.GET("/order/list", orderHandler.GetOrderList)

	// Order reports
	auth.GET("/order/:id/invoice", reportHandler.GetOrderInvoice)

	// Upload
	auth.POST("/upload", uploadHandler.UploadFile)
	auth.POST("/upload/batch", uploadHandler.BatchUploadFiles)
	auth.POST("/upload/delete", uploadHandler.DeleteFile)

	// Static files for uploads
	s.router.Static("/uploads", "./uploads")
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
