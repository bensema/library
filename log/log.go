package log

import (
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(format, args...))
}

func Infof(format string, args ...interface{}) {
	logger.Info(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	logger.Warn(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	logger.Error(fmt.Sprintf(format, args...))
}
