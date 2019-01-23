package console

import (
	"github.com/shanexu/logn/appender/encoder"
	ec "github.com/shanexu/logn/appender/encoder/common"
	"github.com/shanexu/logn/common"
	"go.uber.org/zap/zapcore"
)

var defaultConfig = ec.JsonEncoderConfig{
	TimeKey:       "ts",
	LevelKey:      "level",
	NameKey:       "logger",
	CallerKey:     "caller",
	MessageKey:    "msg",
	StacktraceKey: "stacktrace",
	LineEnding:    "\n",
	TimeEncoder:   "epoch",
}

func init() {
	encoder.RegisterType("console", func(cfg *common.Config) (encoder.Encoder, error) {
		config := defaultConfig
		if cfg != nil {
			if err := cfg.Unpack(&config); err != nil {
				return nil, err
			}
		}

		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        config.TimeKey,
			LevelKey:       config.LevelKey,
			NameKey:        config.NameKey,
			CallerKey:      config.CallerKey,
			MessageKey:     config.MessageKey,
			StacktraceKey:  config.StacktraceKey,
			LineEnding:     config.LineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		te, err := ec.GetTimeEncoder(config.TimeEncoder)
		if err != nil {
			return nil, err
		}
		encoderConfig.EncodeTime = te

		return zapcore.NewConsoleEncoder(encoderConfig), nil
	})
}
