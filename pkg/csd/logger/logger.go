package logger

import (
	"github.com/ArenDjango/golang-test-task/pkg/csd/logger/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogManager struct {
	cfg Config
}

func (l *LogManager) InitLogger() {
	log.Structured = l.cfg.Structured

	var logConfig zap.Config
	if l.cfg.Structured {
		logConfig = zap.Config{
			Level:       zap.NewAtomicLevelAt(zapcore.Level(loggerLevelMap[l.cfg.Level])),
			Development: true,
			Encoding:    "json",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}
	} else {
		logConfig = zap.Config{
			Level:       zap.NewAtomicLevelAt(zapcore.Level(loggerLevelMap[l.cfg.Level])),
			Development: true,
			Encoding:    "console",
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "T",
				LevelKey:       "L",
				NameKey:        "N",
				CallerKey:      "C",
				MessageKey:     "M",
				StacktraceKey:  "S",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalColorLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}
	}

	var err error
	Logger, err := logConfig.Build(zap.AddCallerSkip(l.cfg.SkipFrameCount))
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(Logger)
	Logger.Info("logger started with settings", zap.Any("config", l.cfg))
}
