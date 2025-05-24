package database

import "fmt"

// DatabaseConfigBuilder 数据库配置构建器
type DatabaseConfigBuilder struct {
	config *DatabaseConfig
}

// NewDatabaseConfigBuilder 创建配置构建器
func NewDatabaseConfigBuilder() *DatabaseConfigBuilder {
	return &DatabaseConfigBuilder{
		config: DefaultDatabaseConfig(),
	}
}

// Host 设置主机
func (b *DatabaseConfigBuilder) Host(host string) *DatabaseConfigBuilder {
	b.config.Host = host
	return b
}

// Port 设置端口
func (b *DatabaseConfigBuilder) Port(port string) *DatabaseConfigBuilder {
	b.config.Port = port
	return b
}

// Credentials 设置用户名和密码
func (b *DatabaseConfigBuilder) Credentials(username, password string) *DatabaseConfigBuilder {
	b.config.Username = username
	b.config.Password = password
	return b
}

// Database 设置数据库名
func (b *DatabaseConfigBuilder) Database(dbname string) *DatabaseConfigBuilder {
	b.config.DBName = dbname
	return b
}

// ConnectionPool 设置连接池
func (b *DatabaseConfigBuilder) ConnectionPool(maxOpen, maxIdle int) *DatabaseConfigBuilder {
	b.config.MaxOpenConns = maxOpen
	b.config.MaxIdleConns = maxIdle
	return b
}

// Build 构建配置
func (b *DatabaseConfigBuilder) Build() (*DatabaseConfig, error) {
	if err := b.config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid database config: %w", err)
	}
	return b.config.Clone(), nil
}
