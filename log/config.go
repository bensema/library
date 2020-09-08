package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Config struct {
	Level string

	Debug bool

	LevelSeparate bool

	// stdout
	Stdout bool

	// lumberjack config
	Lbj *LumberjackConfig

	// log-agent
	Agent *AgentConfig
}

func Init(c *Config) {

	var cores []zapcore.Core
	if c.Debug {
		cores = append(cores, newCore(c, zap.DebugLevel))
	}
	if c.LevelSeparate {
		cores = append(cores, newCore(c, zap.InfoLevel))
		cores = append(cores, newCore(c, zap.WarnLevel))
		cores = append(cores, newCore(c, zap.ErrorLevel))
	} else {
		cores = append(cores, newCore(c, zap.InfoLevel))

	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "admin-service"))
	// 构造日志
	logger = zap.New(zapcore.NewTee(cores...), caller, development, filed)

	//logger.Info("log 初始化成功")
	//logger.Info("无",
	//	zap.String("url", "http://www.baidu.com"),
	//	zap.Int("attempt", 3),
	//	zap.Duration("backoff", time.Second))
}

func newCore(c *Config, l zapcore.Level) zapcore.Core {
	var (
		name string
		wss  []zapcore.WriteSyncer
	)

	switch l {
	case zapcore.DebugLevel:
		name = "debug"
	case zapcore.InfoLevel:
		name = "info"
	case zapcore.WarnLevel:
		name = "warn"
	case zapcore.ErrorLevel:
		name = "error"
	case zapcore.FatalLevel:
		name = "fatal"
	default:
		name = "log"
	}

	hook := NewLbj(c.Lbj, name)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	wss = append(wss, zapcore.AddSync(hook))

	if c.Stdout {
		wss = append(wss, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		zapcore.NewMultiWriteSyncer(wss...),   // 打印到控制台和文件
		l,                                     // 日志级别
	)
	return core
}
