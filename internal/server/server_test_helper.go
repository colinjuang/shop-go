package server

import (
	"github.com/colinjuang/shop-go/internal/config"
	"github.com/colinjuang/shop-go/internal/pkg/minio"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"gorm.io/gorm"
)

// MockServerContext 创建用于测试的模拟服务器上下文
type MockServerContext struct {
	*ServerContext
}

// NewMockServerContext 创建测试用的模拟服务器
func NewMockServerContext(db *gorm.DB, redisClient *redis.Client, minioClient *minio.Client) *MockServerContext {
	mockServer := &ServerContext{
		config: &config.Config{}, // 使用默认配置
		DB:     db,
		Redis:  redisClient,
		Minio:  minioClient,
	}

	// 设置为全局实例（仅用于测试）
	SetServer(mockServer)

	return &MockServerContext{
		ServerContext: mockServer,
	}
}

// TestHelper 测试辅助结构
type TestHelper struct {
	originalServer *ServerContext
}

// SetupTest 设置测试环境
func (h *TestHelper) SetupTest(mockDB *gorm.DB, mockRedis *redis.Client, mockMinio *minio.Client) {
	// 保存原始的服务器实例
	h.originalServer = GetServer()

	// 设置模拟实例
	NewMockServerContext(mockDB, mockRedis, mockMinio)
}

// TeardownTest 清理测试环境
func (h *TestHelper) TeardownTest() {
	// 恢复原始实例
	if h.originalServer != nil {
		SetServer(h.originalServer)
	} else {
		ResetServer()
	}
}

// NewTestHelper 创建测试辅助器
func NewTestHelper() *TestHelper {
	return &TestHelper{}
}
