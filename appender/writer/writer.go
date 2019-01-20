package writer

import "go.uber.org/zap/zapcore"

type Writer interface {
	zapcore.WriteSyncer
}