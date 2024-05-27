package logger

import (
	"github.com/ArenDjango/golang-test-task/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

// Logger is a logger :)
type Logger struct {
	*zap.Logger
}

var (
	logger Logger
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Logger {
	once.Do(func() {
		//zapLogger := logger.Sync()
		cfg := config.Get()
		zapLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
		//// Set proper loglevel based on config
		switch cfg.LogLevel {
		case "debug":
			zapLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		case "info":
			zapLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		case "warn", "warning":
			zapLevel = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		case "err", "error":
			zapLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		case "fatal":
			zapLevel = zap.NewAtomicLevelAt(zapcore.FatalLevel)
		case "panic":
			zapLevel = zap.NewAtomicLevelAt(zapcore.PanicLevel)
		default:
			zapLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		}
		//logger = Logger{&zeroLogger}
		//
		cfgZap := zap.Config{
			Level:            zapLevel, // Уровень логирования
			Development:      true,
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		// Создание логгера на основе конфигурации
		logger, err := cfgZap.Build()
		if err != nil {
			panic(err)
		}
		defer logger.Sync() // flushes buffer, if any

		// Установка глобального логгера
		zap.ReplaceGlobals(logger)
	})
	return &logger
}
