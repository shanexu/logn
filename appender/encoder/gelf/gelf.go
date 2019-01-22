package gelf

import (
	"github.com/shanexu/logn/appender/encoder"
	"github.com/shanexu/logn/common"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
)

type Encoder struct {
	fields []zapcore.Field
	zapcore.Encoder
}

func (e *Encoder) EncodeEntry(enc zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	newFields := make([]zap.Field, len(e.fields)+len(fields))
	i := 0
	for ; i < len(e.fields); i++ {
		newFields[i] = e.fields[i]
	}
	for ; i < len(e.fields)+len(fields); i++ {
		j := i - len(e.fields)
		f := fields[j]
		f.Key = "_" + f.Key
		newFields[i] = f
	}
	return e.Encoder.EncodeEntry(enc, newFields)
}

func LevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	level := uint8(7)
	switch l {
	case zapcore.DebugLevel:
		level = 7
	case zapcore.InfoLevel:
		level = 6
	case zapcore.WarnLevel:
		level = 4
	case zapcore.ErrorLevel:
		level = 3
	case zapcore.DPanicLevel:
		level = 2
	case zapcore.PanicLevel:
		level = 1
	case zapcore.FatalLevel:
		level = 0
	}
	enc.AppendUint8(level)
}

func init() {
	encoder.RegisterType("gelf", func(config *common.Config) (i encoder.Encoder, e error) {
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "_logger",
			CallerKey:      "_caller",
			MessageKey:     "short_message",
			StacktraceKey:  "full_message",
			LineEnding:     "\n",
			EncodeLevel:    LevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		hostname, err := os.Hostname()
		if err != nil {
			return nil, err
		}
		fields := []zapcore.Field{
			zap.String("version", "1.1"),
			zap.String("host", hostname),
		}
		return &Encoder{
			fields:  fields,
			Encoder: zapcore.NewJSONEncoder(encoderConfig),
		}, nil
	})
}
