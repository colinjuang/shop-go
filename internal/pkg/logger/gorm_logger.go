package logger

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormLogger is a custom logger for GORM
type GormLogger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  logger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

// NewGormLogger creates a new GORM logger that uses zap
func NewGormLogger() *GormLogger {
	return &GormLogger{
		ZapLogger:                 log,
		LogLevel:                  logger.Info,
		SlowThreshold:             200 * time.Millisecond,
		IgnoreRecordNotFoundError: true,
	}
}

// LogMode sets the log level
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info logs info messages
func (l *GormLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel >= logger.Info {
		Sugar.Infof(msg, args...)
	}
}

// Warn logs warn messages
func (l *GormLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel >= logger.Warn {
		Sugar.Warnf(msg, args...)
	}
}

// Error logs error messages
func (l *GormLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel >= logger.Error {
		Sugar.Errorf(msg, args...)
	}
}

// Trace logs SQL and time
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	// Log based on error and execution time
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		l.ZapLogger.Error("GORM Error",
			zap.Error(err),
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows))
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		l.ZapLogger.Warn("GORM Slow Query",
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows))
	case l.LogLevel >= logger.Info:
		l.ZapLogger.Debug("GORM Query",
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows))
	}
}
