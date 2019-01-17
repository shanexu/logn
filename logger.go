package logn

import (
	"go.uber.org/zap"
)

type Logger interface {

	// Debug uses fmt.Sprint to construct and log a message.
	Debug(args ...interface{})

	// Info uses fmt.Sprint to construct and log a message.
	Info(args ...interface{})

	// Warn uses fmt.Sprint to construct and log a message.
	Warn(args ...interface{})

	// Error uses fmt.Sprint to construct and log a message.
	Error(args ...interface{})

	// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit(1).
	Fatal(args ...interface{})

	// Panic uses fmt.Sprint to construct and log a message, then panics.
	Panic(args ...interface{})

	// DPanic uses fmt.Sprint to construct and log a message. In development, the
	// logger then panics.
	DPanic(args ...interface{})

	// Debugf uses fmt.Sprintf to construct and log a message.
	Debugf(format string, args ...interface{})

	// Infof uses fmt.Sprintf to log a templated message.
	Infof(format string, args ...interface{})

	// Warnf uses fmt.Sprintf to log a templated message.
	Warnf(format string, args ...interface{})

	// Errorf uses fmt.Sprintf to log a templated message.
	Errorf(format string, args ...interface{})

	// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit(1).
	Fatalf(format string, args ...interface{})

	// Panicf uses fmt.Sprintf to log a templated message, then panics.
	Panicf(format string, args ...interface{})

	// DPanicf uses fmt.Sprintf to log a templated message. In development, the
	// logger then panics.
	DPanicf(format string, args ...interface{})

	// Debugw logs a message with some additional context. The additional context
	// is added in the form of key-value pairs. The optimal way to write the value
	// to the log message will be inferred by the value's type. To explicitly
	// specify a type you can pass a Field such as logp.Stringer.
	Debugw(msg string, keysAndValues ...interface{})

	// Infow logs a message with some additional context. The additional context
	// is added in the form of key-value pairs. The optimal way to write the value
	// to the log message will be inferred by the value's type. To explicitly
	// specify a type you can pass a Field such as logp.Stringer.
	Infow(msg string, keysAndValues ...interface{})

	// Warnw logs a message with some additional context. The additional context
	// is added in the form of key-value pairs. The optimal way to write the value
	// to the log message will be inferred by the value's type. To explicitly
	// specify a type you can pass a Field such as logp.Stringer.
	Warnw(msg string, keysAndValues ...interface{})

	// Errorw logs a message with some additional context. The additional context
	// is added in the form of key-value pairs. The optimal way to write the value
	// to the log message will be inferred by the value's type. To explicitly
	// specify a type you can pass a Field such as logp.Stringer.
	Errorw(msg string, keysAndValues ...interface{})

	// Fatalw logs a message with some additional context, then calls os.Exit(1).
	// The additional context is added in the form of key-value pairs. The optimal
	// way to write the value to the log message will be inferred by the value's
	// type. To explicitly specify a type you can pass a Field such as
	// logp.Stringer.
	Fatalw(msg string, keysAndValues ...interface{})

	// Panicw logs a message with some additional context, then panics. The
	// additional context is added in the form of key-value pairs. The optimal way
	// to write the value to the log message will be inferred by the value's type.
	// To explicitly specify a type you can pass a Field such as logp.Stringer.
	Panicw(msg string, keysAndValues ...interface{})

	// DPanicw logs a message with some additional context. The logger panics only
	// in Development mode.  The additional context is added in the form of
	// key-value pairs. The optimal way to write the value to the log message will
	// be inferred by the value's type. To explicitly specify a type you can pass a
	// Field such as logp.Stringer.
	DPanicw(msg string, keysAndValues ...interface{})
}

type ZapLogger struct {
	sugar *zap.SugaredLogger
}

func (z *ZapLogger) Debug(args ...interface{}) {
	z.sugar.Debug(args)
}

func (z *ZapLogger) Info(args ...interface{}) {
	z.sugar.Info(args)
}

func (z *ZapLogger) Warn(args ...interface{}) {
	z.sugar.Warn(args)
}

func (z *ZapLogger) Error(args ...interface{}) {
	z.sugar.Error(args)
}

func (z *ZapLogger) Fatal(args ...interface{}) {
	z.sugar.Fatal(args)
}

func (z *ZapLogger) Panic(args ...interface{}) {
	z.sugar.Panic(args)
}

func (z *ZapLogger) DPanic(args ...interface{}) {
	z.sugar.DPanic(args)
}

func (z *ZapLogger) Debugf(format string, args ...interface{}) {
	z.sugar.Debugf(format, args)
}

func (z *ZapLogger) Infof(format string, args ...interface{}) {
	z.sugar.Infof(format, args)
}

func (z *ZapLogger) Warnf(format string, args ...interface{}) {
	z.sugar.Warnf(format, args)
}

func (z *ZapLogger) Errorf(format string, args ...interface{}) {
	z.sugar.Errorf(format, args)
}

func (z *ZapLogger) Fatalf(format string, args ...interface{}) {
	z.sugar.Fatalf(format, args)
}

func (z *ZapLogger) Panicf(format string, args ...interface{}) {
	z.sugar.Panicf(format, args)
}

func (z *ZapLogger) DPanicf(format string, args ...interface{}) {
	z.sugar.DPanicf(format, args)
}

func (z *ZapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	z.sugar.Debugw(msg, keysAndValues)
}

func (z *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	z.sugar.Infow(msg, keysAndValues)
}

func (z *ZapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	z.sugar.Warnw(msg, keysAndValues)
}

func (z *ZapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	z.sugar.Errorw(msg, keysAndValues)
}

func (z *ZapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	z.sugar.Fatalw(msg, keysAndValues)
}

func (z *ZapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	z.sugar.Panicw(msg, keysAndValues)
}

func (z *ZapLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	z.sugar.DPanicw(msg, keysAndValues)
}
