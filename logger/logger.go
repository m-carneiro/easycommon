package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ProductionEnvName = "production"
)

var (
	logger Logger
)

func Initialize(enviroment string) {
	var config zap.Config
	config = NewProductionConfig()
	if enviroment != ProductionEnvName {
		config = zap.NewDevelopmentConfig()
	}

	config.DisableStacktrace = true
	log, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger = log.WithOptions(zap.AddCallerSkip(1)).Sugar()
}

func NewProductionConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},

		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
	}
}

func NewProductionEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func WithLogger(l Logger) {
	logger = l
}

func Debug(msg ...interface{}) {
	logger.Debug(msg...)
}

func Debugf(format string, msg ...interface{}) {
	logger.Debugf(format, msg...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

func Info(msg ...interface{}) {
	logger.Info(msg...)
}

func Infof(format string, msg ...interface{}) {
	logger.Infof(format, msg...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

func Warn(msg ...interface{}) {
	logger.Warn(msg...)
}

func Warnf(format string, msg ...interface{}) {
	logger.Warnf(format, msg...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues...)
}

func Error(msg ...interface{}) {
	logger.Error(msg...)
}

func Errorf(format string, msg ...interface{}) {
	logger.Errorf(format, msg...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues...)
}

func Fatal(msg ...interface{}) {
	logger.Fatal(msg...)
}

func Fatalf(format string, msg ...interface{}) {
	logger.Fatalf(format, msg...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Fatalw(msg, keysAndValues...)
}

func Panic(msg ...interface{}) {
	logger.Panic(msg...)
}

func Panicf(format string, msg ...interface{}) {
	logger.Panicf(format, msg...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panicw(msg, keysAndValues...)
}
