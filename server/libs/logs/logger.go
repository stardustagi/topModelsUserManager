package logs

import (
	"logistics/utils"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

func Init(logConfigJson []byte) {
	// * lumberjack.Logger 用于日志轮转
	var logConfig *lumberjack.Logger
	var err error
	if logConfig, err = utils.Bytes2Struct[*lumberjack.Logger](logConfigJson); err != nil {
		panic("Failed to parse log configuration: " + err.Error())
	}
	// 日志级别
	var encoderCfg zapcore.EncoderConfig
	level := zapcore.InfoLevel

	// 编码器配置
	encoderCfg = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var zapCore []zapcore.Core
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	// 控制台输出
	consoleWriter := zapcore.Lock(os.Stdout)
	zapCore = append(zapCore, zapcore.NewCore(
		encoder,
		consoleWriter,
		level,
	))
	// 文件输出配置
	var fileWriter zapcore.WriteSyncer
	if logConfig == nil {
		fileWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    500, // MB
			MaxBackups: 5,
			MaxAge:     7, // days
			Compress:   true,
		})
	} else {
		fileWriter = zapcore.AddSync(logConfig)
	}

	zapCore = append(zapCore, zapcore.NewCore(
		encoder,
		fileWriter,
		level,
	))

	// 编码器统一使用 JSON

	// 合并两个输出目标
	core := zapcore.NewTee(zapCore...)

	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Infof(format string, args ...interface{}) {
	if Log != nil {
		Log.Sugar().Infof(format, args...)
	}
}

func Info(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Info(msg, fields...)
	}
}

func Warnf(format string, args ...interface{}) {
	if Log != nil {
		Log.Sugar().Warnf(format, args...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Warn(msg, fields...)
	}
}

func Errorf(format string, args ...interface{}) {
	if Log != nil {
		Log.Sugar().Errorf(format, args...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Error(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	if Log != nil {
		Log.Error(msg, fields...)
	}
}

func Debugf(format string, args ...interface{}) {
	if Log != nil {
		Log.Sugar().Errorf(format, args)
	}
}
