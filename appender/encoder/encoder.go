package encoder

import "go.uber.org/zap/zapcore"

type Encoder interface {
	zapcore.Encoder
}