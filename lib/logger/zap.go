package logger

import (
	"fmt"
	"go-app/config"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zapLogger *zap.Logger

func init() {
	// 设置日志配置格式
	encoderConf := zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		MessageKey:   "msg",
		CallerKey:    "file",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}

	hook := lumberjack.Logger{
		Filename:   strings.TrimRight(config.LOG.Path, "/") + "/" + config.LOG.File,
		MaxSize:    config.LOG.MaxSize,
		MaxBackups: config.LOG.MaxBackup,
		MaxAge:     config.LOG.MaxAge,
		Compress:   config.LOG.Compress,
	}

	// 设置日志级别
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(config.LOG.Level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 Level 配置项")
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(*logLevel)

	var (
		writer  zapcore.WriteSyncer
		encoder zapcore.Encoder
	)

	if config.LOG.Type == "stdout" {
		writer = zapcore.AddSync(os.Stdout)
		encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConf)
	} else {
		if config.APP.Debug {
			writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
			encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
			encoder = zapcore.NewConsoleEncoder(encoderConf)
		} else {
			writer = zapcore.AddSync(&hook)
			encoder = zapcore.NewJSONEncoder(encoderConf)
		}
	}

	// 初始化
	core := zapcore.NewCore(encoder, writer, logLevel)

	zapLogger = zap.New(
		core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	zap.ReplaceGlobals(zapLogger)

	defer zapLogger.Sync()
}

func Debug(msg string, fields ...zapcore.Field) {
	zapLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	zapLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	zapLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	zapLogger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zapcore.Field) {
	zapLogger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	zapLogger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	zapLogger.Fatal(msg, fields...)
}

func Suger() *zap.SugaredLogger {
	return zapLogger.Sugar()
}
