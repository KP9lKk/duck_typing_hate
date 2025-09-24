package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Logger *zap.Logger
}

func New(level string) *Logger {
	var l zapcore.Level

	switch strings.ToLower(level) {
	case "error":
		l = zapcore.ErrorLevel
	case "warn":
		l = zapcore.WarnLevel
	case "info":
		l = zapcore.InfoLevel
	case "debug":
		l = zapcore.DebugLevel
	default:
		l = zapcore.InfoLevel
	}
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))

	atom.SetLevel(l)
	return &Logger{Logger: logger}
}
