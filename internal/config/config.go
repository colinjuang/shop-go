package config

import (
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Server       ServerConfig            `mapstructure:"server"`
	DatabaseConf database.DatabaseConfig `mapstructure:"database"`
	Redis        RedisConfig             `mapstructure:"redis"`
	MinIO        MinIOConfig             `mapstructure:"minio"`
	JWT          JWTConfig               `mapstructure:"jwt"`
	Wechat       WechatConfig            `mapstructure:"wechat"`
	Upload       UploadConfig            `mapstructure:"upload"`
	Logger       LoggerConfig            `mapstructure:"logger"`
}

// LoggerConfig represents logger configuration
type LoggerConfig struct {
	Level      string `mapstructure:"level"`
	Encoding   string `mapstructure:"encoding"`
	OutputPath string `mapstructure:"output_path"`
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Port         string `mapstructure:"port"`
	Environment  string `mapstructure:"environment"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	IdleTimeout  int    `mapstructure:"idle_timeout"`
}

// DatabaseConfig represents database configuration
// type DatabaseConfig struct {
// 	Host     string `mapstructure:"host"`
// 	Port     string `mapstructure:"port"`
// 	Username string `mapstructure:"username"`
// 	Password string `mapstructure:"password"`
// 	DBName   string `mapstructure:"dbname"`

// 	// 连接池配置（可选，有默认值）
// 	MaxOpenConns int `mapstructure:"max_open_conns"`
// 	MaxIdleConns int `mapstructure:"max_idle_conns"`

// 	// 高级配置（可选）
// 	LogLevel      string `mapstructure:"log_level"`
// 	TablePrefix   string `mapstructure:"table_prefix"`
// 	SingularTable *bool  `mapstructure:"singular_table"` // 使用指针以支持默认值
// }

// RedisConfig represents Redis configuration
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	Prefix   string `mapstructure:"prefix"`
}

// MinIOConfig represents MinIO configuration
type MinIOConfig struct {
	Endpoint  string `mapstructure:"endpoint"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	UseSSL    bool   `mapstructure:"use_ssl"`
	Bucket    string `mapstructure:"bucket"`
	Location  string `mapstructure:"location"`
}

// JWTConfig represents JWT configuration
type JWTConfig struct {
	Secret    string `mapstructure:"secret"`
	ExpiresIn int    `mapstructure:"expires_in"` // in hours
}

// WechatConfig represents WeChat configuration
type WechatConfig struct {
	AppID     string `mapstructure:"app_id"`
	AppSecret string `mapstructure:"app_secret"`
}

// UploadConfig represents file upload configuration
type UploadConfig struct {
	SavePath string `mapstructure:"save_path"`
	MaxSize  int64  `mapstructure:"max_size"` // in bytes
}

// LoadConfig loads configuration from config file
func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// GetConfig gets the application configuration
func GetConfig() *Config {
	config, err := LoadConfig("configs/config.yaml")
	if err == nil {
		return config
	}

	// 所有路径都加载失败
	panic("Failed to load configuration: " + err.Error())
}
