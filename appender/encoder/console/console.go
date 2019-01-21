package console

import (
	"github.com/shanexu/logn/appender/encoder"
	"github.com/shanexu/logn/common"
	"go.uber.org/zap/zapcore"
)

// Config is used to pass encoding parameters to New.
type Config struct {
	TimeKey       string `json:"time_key"`
	LevelKey      string `json:"level_key"`
	NameKey       string `json:"name_key"`
	CallerKey     string `json:"caller_key"`
	MessageKey    string `json:"message_key"`
	StacktraceKey string `json:"stacktrace_key"`
	LineEnding    string `json:"line_ending"`
}

var defaultConfig = Config{
	TimeKey:       "ts",
	LevelKey:      "level",
	NameKey:       "logger",
	CallerKey:     "caller",
	MessageKey:    "msg",
	StacktraceKey: "stacktrace",
	LineEnding:    "\n",
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
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		return zapcore.NewConsoleEncoder(encoderConfig), nil
	})
}

