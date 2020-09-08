package log

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LumberjackConfig struct {
	Dir string

	MaxSize int

	MaxBackups int

	MaxAge int

	Compress bool
}

func NewLbj(c *LumberjackConfig, name string) *lumberjack.Logger {
	lbj := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", c.Dir, name), // 日志文件路径
		MaxSize:    c.MaxSize,                             // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: c.MaxBackups,                          // 日志文件最多保存多少个备份
		MaxAge:     c.MaxAge,                              // 文件最多保存多少天
		Compress:   c.Compress,                            // 是否压缩
	}
	return lbj
}
