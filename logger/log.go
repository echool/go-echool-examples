package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	Log = initLogger(2, zap.DebugLevel, "./logs/log.log", 10, 30, 30, true)
}

func initLogger(
	logType uint, // 0:console; 1:file; 2:console&file
	logLevel zapcore.Level,
	logPath string,
	maxSize int,
	maxAge int,
	maxBackups int,
	compress bool,
) *zap.Logger {
	if logType > 2 {
		panic("logType must be between 0-2")
	}

	stdWriter, fileWriter := zapcore.Lock(os.Stdout), zapcore.Lock(os.Stdout)

	if logType >= 1 {
		fileWriter = zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   logPath,
				MaxSize:    maxSize,
				MaxBackups: maxBackups,
				MaxAge:     maxAge,
				Compress:   compress,
				LocalTime:  true,
			})
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	consuleEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return logType != 1 && lvl >= logLevel
	})
	fileEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return logType >= 1 && lvl >= logLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdWriter, consuleEnabler),
		zapcore.NewCore(fileEncoder, fileWriter, fileEnabler),
	)

	return zap.New(core)
}
