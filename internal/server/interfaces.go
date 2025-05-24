package server

import "github.com/gin-gonic/gin"

// HTTPServer defines the interface for HTTP server
type HTTPServer interface {
	Start() error
	GetRouter() *gin.Engine
	Shutdown() error
}

// RouterRegistrar defines the interface for route registration
type RouterRegistrar interface {
	RegisterRoutes(router *gin.Engine)
}
