package appender

import "go.uber.org/zap/zapcore"

type Appender struct {
	Writer  zapcore.WriteSyncer
	Encoder zapcore.Encoder
}
