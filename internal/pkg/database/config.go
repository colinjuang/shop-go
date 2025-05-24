package database

import (
	"errors"
	"fmt"
	"time"
)

// DatabaseConfig 数据库配置结构体
type DatabaseConfig struct {
	// 连接配置
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`

	// 连接池配置
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`

	// 其他配置
	LogLevel        string `mapstructure:"log_level"`
	EnableMigration bool   `mapstructure:"enable_migration"`

	// 高级配置
	TablePrefix   string        `mapstructure:"table_prefix"`
	SingularTable bool          `mapstructure:"singular_table"`
	Timeout       time.Duration `mapstructure:"timeout"`
}

// DefaultDatabaseConfig 返回默认数据库配置
func DefaultDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:            "localhost",
		Port:            "3306",
		MaxOpenConns:    100,
		MaxIdleConns:    10,
		ConnMaxLifetime: time.Hour,
		ConnMaxIdleTime: 30 * time.Minute,
		LogLevel:        "info",
		EnableMigration: false,
		TablePrefix:     "",
		SingularTable:   true,
		Timeout:         30 * time.Second,
	}
}

// Validate 验证数据库配置并设置默认值
func (cfg *DatabaseConfig) Validate() error {
	if cfg.Host == "" {
		return errors.New("database host is required")
	}
	if cfg.Port == "" {
		return errors.New("database port is required")
	}
	if cfg.Username == "" {
		return errors.New("database username is required")
	}
	if cfg.DBName == "" {
		return errors.New("database name is required")
	}

	// 设置默认值（如果为0）
	if cfg.MaxOpenConns <= 0 {
		cfg.MaxOpenConns = 100
	}
	if cfg.MaxIdleConns <= 0 {
		cfg.MaxIdleConns = 10
	}
	if cfg.ConnMaxLifetime <= 0 {
		cfg.ConnMaxLifetime = time.Hour
	}
	if cfg.ConnMaxIdleTime <= 0 {
		cfg.ConnMaxIdleTime = 30 * time.Minute
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 30 * time.Second
	}

	// 验证连接池配置
	if cfg.MaxIdleConns > cfg.MaxOpenConns {
		return errors.New("max_idle_conns cannot be greater than max_open_conns")
	}

	return nil
}

// DSN 生成数据库连接字符串
func (cfg *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}

// String 返回配置的字符串表示（隐藏敏感信息）
func (cfg *DatabaseConfig) String() string {
	return fmt.Sprintf("Database{Host: %s, Port: %s, DBName: %s, MaxOpenConns: %d, MaxIdleConns: %d}",
		cfg.Host, cfg.Port, cfg.DBName, cfg.MaxOpenConns, cfg.MaxIdleConns)
}

// Clone 创建配置的副本
func (cfg *DatabaseConfig) Clone() *DatabaseConfig {
	clone := *cfg
	return &clone
}

// WithMaxOpenConns 设置最大连接数（链式调用）
func (cfg *DatabaseConfig) WithMaxOpenConns(max int) *DatabaseConfig {
	cfg.MaxOpenConns = max
	return cfg
}

// WithMaxIdleConns 设置最大空闲连接数（链式调用）
func (cfg *DatabaseConfig) WithMaxIdleConns(max int) *DatabaseConfig {
	cfg.MaxIdleConns = max
	return cfg
}

// WithLogLevel 设置日志级别（链式调用）
func (cfg *DatabaseConfig) WithLogLevel(level string) *DatabaseConfig {
	cfg.LogLevel = level
	return cfg
}

// WithTimeout 设置超时时间（链式调用）
func (cfg *DatabaseConfig) WithTimeout(timeout time.Duration) *DatabaseConfig {
	cfg.Timeout = timeout
	return cfg
}
