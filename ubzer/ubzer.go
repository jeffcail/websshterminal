package ubzer

import (
	"os"

	_const "github.com/c/websshterminal.io/const"

	"github.com/robfig/cron"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var MLog *zap.Logger

// InitLogger
func InitLogger(filepath string) {
	encoder := getEncoder()
	writeSyncer := getLogWriter(filepath)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	consoleDebug := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	p := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})
	var allCode []zapcore.Core
	allCode = append(allCode, core)
	allCode = append(allCode, zapcore.NewCore(consoleEncoder, consoleDebug, p))
	c := zapcore.NewTee(allCode...)
	MLog = zap.New(c, zap.AddCaller())
}

// getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(_const.Layout)
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter
func getLogWriter(path string) zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:  path,
		MaxSize:   10240,
		MaxAge:    7,
		LocalTime: true,
		Compress:  true,
	}
	c := cron.New()
	c.AddFunc("0 0 0 1/1 * ?", func() {
		lumberjackLogger.Rotate()
	})
	c.Start()
	return zapcore.AddSync(lumberjackLogger)
}

var EchoLog *EchoLogger

type EchoLogger struct{}

func (this *EchoLogger) Write(p []byte) (n int, err error) {
	MLog.Info("ECHO", zap.String("请求", string(p)))
	return len(p), nil
}
