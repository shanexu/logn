package common

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

// Config is used to pass encoding parameters to New.
type JsonEncoderConfig struct {
	TimeKey       string `logn-config:"time_key"`
	LevelKey      string `logn-config:"level_key"`
	NameKey       string `logn-config:"name_key"`
	CallerKey     string `logn-config:"caller_key"`
	MessageKey    string `logn-config:"message_key"`
	StacktraceKey string `logn-config:"stacktrace_key"`
	LineEnding    string `logn-config:"line_ending"`
	TimeEncoder   string `logn-config:"time_encoder" logn-validate:"logn.oneof=epoch epoch_millis epoch_nanos ISO8601"`
}

func GetTimeEncoder(name string) (zapcore.TimeEncoder, error) {
	switch name {
	case "epoch":
		return zapcore.EpochTimeEncoder, nil
	case "epoch_millis":
		return zapcore.EpochMillisTimeEncoder, nil
	case "epoch_nanos":
		return zapcore.EpochNanosTimeEncoder, nil
	case "ISO8601":
		return zapcore.ISO8601TimeEncoder, nil
	default:
		return nil, fmt.Errorf("no such TimeEncoder %q", name)
	}
}
