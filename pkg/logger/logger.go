package logger

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Field = zapcore.Field

var highPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
})

var lowPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
})

var debugPriority = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl == zapcore.DebugLevel
})

func Logger(logpath, errpath, debugPath string, debug bool) (*zap.Logger, error) {
	jsonEncoder := encoder(debug)

	//由于再次封装日志，因此需要打印上一级的调用，1表示向上跳一级
	callerSkip := zap.AddCallerSkip(1)

	logger, err := zap.NewProduction(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(
			zapcore.NewCore(
				jsonEncoder,
				zapcore.AddSync(rotateLogger(logpath)),
				lowPriority,
			),
			zapcore.NewCore(
				jsonEncoder,
				zapcore.AddSync(rotateLogger(errpath)),
				highPriority,
			),
			zapcore.NewCore(
				jsonEncoder,
				zapcore.AddSync(rotateLogger(debugPath)),
				debugPriority,
			),
		)
	}), callerSkip)
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	return logger, nil
}

// 编码
func encoder(debug bool) zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()

	if debug {
		ec = zap.NewDevelopmentEncoderConfig()
	}

	ec.EncodeTime = zapcore.RFC3339TimeEncoder
	ec.TimeKey = "time"
	return zapcore.NewJSONEncoder(ec)
}

// 轮转日志
func rotateLogger(output string) io.Writer {
	return &lumberjack.Logger{
		Filename:   output,
		MaxSize:    1000,
		MaxBackups: 300,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	}
}

var logger *zap.Logger

func InitDefault(logpath, errpath, debugPath string, debug bool) (err error) {
	logger, err = Logger(logpath, errpath, debugPath, debug)
	return
}

func Error(msg string, err error, fields ...zapcore.Field) {
	info := []zapcore.Field{zap.Error(err)}
	info = append(info, fields...)
	logger.Error(msg, info...)
}

func Debug(msg string, fields ...Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	logger.Info(msg, fields...)
}
func Warn(msg string, fields ...Field) {
	logger.Warn(msg, fields...)
}
func Panic(msg string, fields ...Field) {
	logger.Panic(msg, fields...)
}
