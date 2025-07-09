package config

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger() (*zap.Logger, error) {
	logLevel := zap.InfoLevel
	if os.Getenv("ENVIRONMENT") == "local" {
		logLevel = zap.DebugLevel
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Development:      true,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "time",
			CallerKey:     "caller",
			FunctionKey:   zapcore.OmitKey,
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	baseLogger, err := config.Build()
	if err != nil {
		fmt.Printf("error-creating-logger")
		os.Exit(100)
	}

	logger := baseLogger.Named("Verbatim")

	return logger, nil
}
