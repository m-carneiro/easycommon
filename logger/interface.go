package logger

type Logger interface {
	Debug(msg ...interface{})
	Debugf(format string, msg ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Info(msg ...interface{})
	Infof(format string, msg ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Warn(msg ...interface{})
	Warnf(format string, msg ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	Error(msg ...interface{})
	Errorf(format string, msg ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Fatal(msg ...interface{})
	Fatalf(format string, msg ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	Panic(msg ...interface{})
	Panicf(format string, msg ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
}
