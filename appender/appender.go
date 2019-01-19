package appender

import (
	"go.uber.org/zap/zapcore"
)

type Writer interface {
	zapcore.WriteSyncer
}

type Appender struct {
	Writer  zapcore.WriteSyncer
	Encoder zapcore.Encoder
}
