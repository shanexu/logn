package logn

import (
	"go.uber.org/zap"

	"github.com/shanexu/logn/core"
)

func Debug(args ...interface{}) {
	logncore.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	logncore.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	logncore.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	logncore.Error(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit(1).
func Fatal(args ...interface{}) {
	logncore.Fatal(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	logncore.Panic(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics.
func DPanic(args ...interface{}) {
	logncore.DPanic(args...)
}

// Debugf uses fmt.Sprintf to construct and log a message.
func Debugf(format string, args ...interface{}) {
	logncore.Debugf(format, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(format string, args ...interface{}) {
	logncore.Infof(format, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(format string, args ...interface{}) {
	logncore.Warnf(format, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(format string, args ...interface{}) {
	logncore.Errorf(format, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	logncore.Fatalf(format, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(format string, args ...interface{}) {
	logncore.Panicf(format, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics.
func DPanicf(format string, args ...interface{}) {
	logncore.DPanicf(format, args...)
}

// Debugw logs a message with some additional context. The additional context
// is added in the form of key-value pairs. The optimal way to write the value
// to the log message will be inferred by the value's type. To explicitly
// specify a type you can pass a Field such as logp.Stringer.
func Debugw(msg string, keysAndValues ...interface{}) {
	logncore.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The additional context
// is added in the form of key-value pairs. The optimal way to write the value
// to the log message will be inferred by the value's type. To explicitly
// specify a type you can pass a Field such as logp.Stringer.
func Infow(msg string, keysAndValues ...interface{}) {
	logncore.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The additional context
// is added in the form of key-value pairs. The optimal way to write the value
// to the log message will be inferred by the value's type. To explicitly
// specify a type you can pass a Field such as logp.Stringer.
func Warnw(msg string, keysAndValues ...interface{}) {
	logncore.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The additional context
// is added in the form of key-value pairs. The optimal way to write the value
// to the log message will be inferred by the value's type. To explicitly
// specify a type you can pass a Field such as logp.Stringer.
func Errorw(msg string, keysAndValues ...interface{}) {
	logncore.Errorw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit(1).
// The additional context is added in the form of key-value pairs. The optimal
// way to write the value to the log message will be inferred by the value's
// type. To explicitly specify a type you can pass a Field such as
// logp.Stringer.
func Fatalw(msg string, keysAndValues ...interface{}) {
	logncore.Fatalw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// additional context is added in the form of key-value pairs. The optimal way
// to write the value to the log message will be inferred by the value's type.
// To explicitly specify a type you can pass a Field such as logp.Stringer.
func Panicw(msg string, keysAndValues ...interface{}) {
	logncore.Panicw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. The logger panics only
// in Development mode.  The additional context is added in the form of
// key-value pairs. The optimal way to write the value to the log message will
// be inferred by the value's type. To explicitly specify a type you can pass a
// Field such as logp.Stringer.
func DPanicw(msg string, keysAndValues ...interface{}) {
	logncore.DPanicw(msg, keysAndValues...)
}


func With(keysAndValues ...interface{}) core.Logger {
	return logncore.Global().(*zap.SugaredLogger).
		With(keysAndValues...).
		Desugar().
		WithOptions(zap.AddCallerSkip(0)).
		Sugar()
}