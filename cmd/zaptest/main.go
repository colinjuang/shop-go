package main

import (
	"os"
	"github.com/colinjuang/shop-go/internal/pkg/logger"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 直接初始化日志
	logConfig := &logger.LogConfig{
		Level:      "debug",   // 使用debug级别以显示所有日志
		Encoding:   "console", // 使用控制台编码以便于阅读
		OutputPath: "stdout",  // 输出到标准输出
	}
	logger.Init(logConfig)
	defer logger.Sync()

	// 记录各种级别的日志
	logger.Debug("这是一条调试日志")
	logger.Info("这是一条信息日志")
	logger.Warn("这是一条警告日志")
	logger.Error("这是一条错误日志", zap.String("error_code", "E101"))

	// 记录带格式的日志
	logger.Debugf("当前时间: %s", time.Now().Format(time.RFC3339))
	logger.Infof("程序已运行 %d 毫秒", 100)

	// 记录带结构化字段的日志
	logger.Info("用户登录",
		zap.String("username", "admin"),
		zap.String("ip", "192.168.1.1"),
		zap.Int("user_id", 1001),
	)

	// 记录错误信息
	err := os.ErrNotExist
	logger.Error("文件操作失败",
		zap.Error(err),
		zap.String("file", "config.json"),
		zap.Duration("latency", 230*time.Millisecond),
	)

	// 测试成功信息
	logger.Info("Zap日志库集成测试成功!")
	logger.Infof("测试完成，共记录 %d 条日志", 9)
}
