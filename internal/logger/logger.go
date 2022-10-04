package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log  *zap.Logger
)

func InitializeZapCustomLogger() {
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:          "json",
		EncoderConfig:     zapcore.EncoderConfig{
			LevelKey: "level",
			TimeKey: "time",
			CallerKey: "file",
			MessageKey: "msg",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:       []string{viper.GetString("logger-output-path"), "stdout"},
	}
	Log, _ =  config.Build()
}
