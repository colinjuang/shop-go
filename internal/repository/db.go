package repository

import (
	"fmt"

	"shop-go/internal/config"
	"shop-go/internal/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	var logLevel gormlogger.LogLevel
	if cfg.Server.Environment == "development" {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}

	// Create custom GORM logger that uses zap
	gormLogger := logger.NewGormLogger()
	gormLogger.LogLevel = logLevel

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
