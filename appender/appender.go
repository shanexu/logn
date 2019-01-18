package appender

import (
	"go.uber.org/zap/zapcore"
)

type Appender interface {
	zapcore.WriteSyncer
}
