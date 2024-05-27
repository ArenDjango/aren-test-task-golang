package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"

	"github.com/pkg/errors"
)

var Logger *zap.Logger
var Structured = false

func init() {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
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

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
}

func Debug(msg string) {
	if !Structured {
		Logger.Debug(msg)
		return
	}
	Logger.Debug(msg, zap.Any("structured", struct{}{}))
}

func Debugf(template string, args ...interface{}) {
	if !Structured {
		Logger.Sugar().Debugf(template, args...)
		return
	}
	Logger.Debug(fmt.Sprintf(template, args...), zap.Any("structured", struct{}{}))
}

func Info(msg string) {
	if !Structured {
		Logger.Info(msg)
		return
	}
	Logger.Info(msg, zap.Any("structured", struct{}{}))
}

func Infof(template string, args ...interface{}) {
	if !Structured {
		Logger.Sugar().Infof(template, args...)
		return
	}
	Logger.Info(fmt.Sprintf(template, args...), zap.Any("structured", struct{}{}))
}

func Warn(msg string) {
	if !Structured {
		Logger.Warn(msg)
		return
	}
	Logger.Warn(msg, zap.Any("structured", struct{}{}))
}

func Warnf(template string, args ...interface{}) {
	if !Structured {
		Logger.Sugar().Warnf(template, args...)
		return
	}
	Logger.Warn(fmt.Sprintf(template, args...), zap.Any("structured", struct{}{}))
}

func Error(err error) {
	if !Structured {
		Logger.Error(err.Error())
		return
	}
	Logger.Error(err.Error(), zap.Any("structured", struct{}{}))
}

func ErrorSentry(err error, traceID string) {
	if traceID != "" {
		err = errors.Wrapf(err, "traceID: %s", traceID)
	}
	// sentry.CaptureException(err) // Uncomment this line if you're using Sentry
	if !Structured {
		Logger.Error(err.Error())
		return
	}
	Logger.Error(err.Error(), zap.Any("structured", struct{}{}))
}

func ErrorSentryIgnoreCtx(err error, traceID string) {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return
	}

	if traceID != "" {
		err = errors.Wrapf(err, "traceID: %s", traceID)
	}
	// sentry.CaptureException(err) // Uncomment this line if you're using Sentry
	if !Structured {
		Logger.Error(err.Error())
		return
	}
	Logger.Error(err.Error(), zap.Any("structured", struct{}{}))
}

func Errorf(template string, args ...interface{}) {
	if !Structured {
		Logger.Sugar().Errorf(template, args...)
		return
	}
	Logger.Error(fmt.Sprintf(template, args...), zap.Any("structured", struct{}{}))
}

func Panic(msg string) {
	if !Structured {
		Logger.Panic(msg)
		return
	}
	Logger.Panic(msg, zap.Any("structured", struct{}{}))
}

func Panicf(template string, args ...interface{}) {
	if !Structured {
		Logger.Sugar().Panicf(template, args...)
		return
	}
	Logger.Panic(fmt.Sprintf(template, args...), zap.Any("structured", struct{}{}))
}

func Fatal(msg string) {
	if !Structured {
		Logger.Fatal(msg)
		return
	}
	Logger.Fatal(msg, zap.Any("structured", struct{}{}))
}

func Fatalf(template string, args ...interface{}) {
	if !Structured {
		Logger.Sugar().Fatalf(template, args...)
		return
	}
	Logger.Fatal(fmt.Sprintf(template, args...), zap.Any("structured", struct{}{}))
}

func DatabaseQuery(name string, query string, took time.Duration) {
	if !Structured {
		Logger.Info("database query", zap.String("name", name), zap.String("query", query), zap.Duration("took", took))
		return
	}
	Logger.Info("database query", zap.Any("structured", struct {
		Name  string
		Query string
		Took  string
	}{
		Name:  name,
		Query: query,
		Took:  fmt.Sprintf("%v", took),
	}))
}
