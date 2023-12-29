package logger

import (
	"dkdns/dkFramework/configs"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Logger *zap.Logger

func Init(appConfigs *configs.Config) {
	// 配置日志编码器，使用 ConsoleEncoder
	config := zap.NewProductionConfig()
	// 设置自定义日志编码器
	config.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: "message",
		TimeKey:    "time", // 时间键名
		LevelKey:   "level",
		CallerKey:  "caller",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(time.RFC3339)) // 自定义时间格式
		},
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大写级别
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// 设置日志文件切割条件
	maxSize := int64(appConfigs.Logging.MaxSizeMB) // MB
	hook := &lumberjack.Logger{
		Filename:   appConfigs.Logging.OutputPath,
		MaxSize:    int(maxSize), // MB
		MaxBackups: 3,            // 最多保留的旧日志文件数
		MaxAge:     28,           // 日志文件保留的最长时间（以天为单位）
		Compress:   true,         // 是否压缩旧的日志文件
	}

	// 设置日志输出到文件
	fileWriter := zapcore.AddSync(hook)

	// 使用 ConsoleEncoder 和文件输出
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config.EncoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), fileWriter),
		zap.InfoLevel, // 输出级别可以根据需求调整
	)

	logger := zap.New(core)

	Logger = logger
}

func Info(message string, fields ...zap.Field) {
	Logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	Logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
}
func Println(a ...interface{}) {
	// 使用 fmt.Sprint() 将任意类型的变量转换为字符串
	logMessage := fmt.Sprint(a...)

	// 输出到日志
	Logger.Info(logMessage)
}
