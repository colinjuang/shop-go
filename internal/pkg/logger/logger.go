package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Global logger instance
	log *zap.Logger
	// Global sugar logger instance
	Sugar *zap.SugaredLogger
	once  sync.Once
)

// LogConfig holds the configuration for the logger
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Encoding   string `mapstructure:"encoding"`
	OutputPath string `mapstructure:"output_path"`
}

// Init initializes the global logger
func Init(cfg *LogConfig) {
	once.Do(func() {
		// Set default values if not provided
		if cfg == nil {
			cfg = &LogConfig{
				Level:      "info",
				Encoding:   "json",
				OutputPath: "stdout",
			}
		}

		// Parse log level
		level := zapcore.InfoLevel
		if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
			level = zapcore.InfoLevel
		}

		// Configure encoder
		var encoder zapcore.Encoder
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.TimeKey = "timestamp"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		if cfg.Encoding == "console" {
			encoder = zapcore.NewConsoleEncoder(encoderConfig)
		} else {
			encoder = zapcore.NewJSONEncoder(encoderConfig)
		}

		// Configure output
		var output zapcore.WriteSyncer
		if cfg.OutputPath == "stdout" {
			output = zapcore.AddSync(os.Stdout)
		} else {
			file, err := os.OpenFile(cfg.OutputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				output = zapcore.AddSync(os.Stdout)
			} else {
				output = zapcore.AddSync(file)
			}
		}

		// Create core
		core := zapcore.NewCore(encoder, output, level)

		// Create logger
		log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
		Sugar = log.Sugar()
	})
}

// Debug logs a message at debug level
func Debug(msg string, fields ...zapcore.Field) {
	ensureLogger()
	log.Debug(msg, fields...)
}

// Info logs a message at info level
func Info(msg string, fields ...zapcore.Field) {
	ensureLogger()
	log.Info(msg, fields...)
}

// Warn logs a message at warn level
func Warn(msg string, fields ...zapcore.Field) {
	ensureLogger()
	log.Warn(msg, fields...)
}

// Error logs a message at error level
func Error(msg string, fields ...zapcore.Field) {
	ensureLogger()
	log.Error(msg, fields...)
}

// Fatal logs a message at fatal level and then calls os.Exit(1)
func Fatal(msg string, fields ...zapcore.Field) {
	ensureLogger()
	log.Fatal(msg, fields...)
}

// Debugf logs a formatted message at debug level
func Debugf(format string, args ...interface{}) {
	ensureLogger()
	Sugar.Debugf(format, args...)
}

// Infof logs a formatted message at info level
func Infof(format string, args ...interface{}) {
	ensureLogger()
	Sugar.Infof(format, args...)
}

// Warnf logs a formatted message at warn level
func Warnf(format string, args ...interface{}) {
	ensureLogger()
	Sugar.Warnf(format, args...)
}

// Errorf logs a formatted message at error level
func Errorf(format string, args ...interface{}) {
	ensureLogger()
	Sugar.Errorf(format, args...)
}

// Fatalf logs a formatted message at fatal level and then calls os.Exit(1)
func Fatalf(format string, args ...interface{}) {
	ensureLogger()
	Sugar.Fatalf(format, args...)
}

// With returns a child logger with the field added to its context
func With(fields ...zapcore.Field) *zap.Logger {
	ensureLogger()
	return log.With(fields...)
}

// Sync flushes any buffered log entries
func Sync() error {
	ensureLogger()
	return log.Sync()
}

// ensureLogger ensures that the logger is initialized
func ensureLogger() {
	if log == nil {
		Init(nil)
	}
}
