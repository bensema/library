package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func specifyLevel(l zapcore.Level) zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == l
	})
}

func level(l zapcore.Level) zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= l
	})
}
