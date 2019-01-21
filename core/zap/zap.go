package zap

import "go.uber.org/zap"

func NewZapLogger(sugar *zap.SugaredLogger) *ZapLogger {
	return &ZapLogger{sugar}
}

type ZapLogger struct {
	sugar *zap.SugaredLogger
}

func (z *ZapLogger) Debug(args ...interface{}) {
	z.sugar.Debug(args...)
}

func (z *ZapLogger) Info(args ...interface{}) {
	z.sugar.Info(args...)
}

func (z *ZapLogger) Warn(args ...interface{}) {
	z.sugar.Warn(args...)
}

func (z *ZapLogger) Error(args ...interface{}) {
	z.sugar.Error(args...)
}

func (z *ZapLogger) Fatal(args ...interface{}) {
	z.sugar.Fatal(args...)
}

func (z *ZapLogger) Panic(args ...interface{}) {
	z.sugar.Panic(args...)
}

func (z *ZapLogger) DPanic(args ...interface{}) {
	z.sugar.DPanic(args...)
}

func (z *ZapLogger) Debugf(format string, args ...interface{}) {
	z.sugar.Debugf(format, args...)
}

func (z *ZapLogger) Infof(format string, args ...interface{}) {
	z.sugar.Infof(format, args...)
}

func (z *ZapLogger) Warnf(format string, args ...interface{}) {
	z.sugar.Warnf(format, args...)
}

func (z *ZapLogger) Errorf(format string, args ...interface{}) {
	z.sugar.Errorf(format, args...)
}

func (z *ZapLogger) Fatalf(format string, args ...interface{}) {
	z.sugar.Fatalf(format, args...)
}

func (z *ZapLogger) Panicf(format string, args ...interface{}) {
	z.sugar.Panicf(format, args...)
}

func (z *ZapLogger) DPanicf(format string, args ...interface{}) {
	z.sugar.DPanicf(format, args...)
}

func (z *ZapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	z.sugar.Debugw(msg, keysAndValues...)
}

func (z *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	z.sugar.Infow(msg, keysAndValues...)
}

func (z *ZapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	z.sugar.Warnw(msg, keysAndValues...)
}

func (z *ZapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	z.sugar.Errorw(msg, keysAndValues...)
}

func (z *ZapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	z.sugar.Fatalw(msg, keysAndValues...)
}

func (z *ZapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	z.sugar.Panicw(msg, keysAndValues...)
}

func (z *ZapLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	z.sugar.DPanicw(msg, keysAndValues...)
}

func (z *ZapLogger) Sync() error {
	return z.sugar.Sync()
}
