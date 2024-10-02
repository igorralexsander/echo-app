package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	*zap.Logger
}

var logInstance *logger

func NewEasyZap() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "msg"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	writerSync := zapcore.Lock(os.Stderr)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writerSync, zapcore.DebugLevel)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	logInstance = &logger{zapLogger}
}

func Info(msg string, fields ...zap.Field) {
	logInstance.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logInstance.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logInstance.Debug(msg, fields...)
}

func Error(err error, msg string, fields ...zap.Field) {
	fields = append(fields, zap.Error(err))
	logInstance.Error(msg, fields...)
}

func Fatal(err error, msg string, fields ...zap.Field) {
	fields = append(fields, zap.Error(err))
	logInstance.Fatal(msg, fields...)
}

func Log(level zapcore.Level, msg string, fields ...zap.Field) {
	logInstance.WithOptions(zap.WithCaller(false)).Log(level, msg, fields...)
}
