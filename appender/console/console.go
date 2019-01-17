package console

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/shanexu/logp/appender"
	"github.com/spf13/viper"
	"os"
)

type Console struct {
	*os.File
}

type Config struct {
	Target `json:"target" validate:"required,oneof=stderr stdout"`
}

type Target = string

const (
	Stdout Target = "stdout"
	Stderr Target = "stderr"
)

var (
	defaultConfig = Config{
		Target: Stdout,
	}
)

func DefaultConfig() Config {
	return defaultConfig
}

func NewConsole(v *viper.Viper) (appender.Appender, error) {
	cfg := DefaultConfig()
	if err := v.Unmarshal(&cfg, func(m *mapstructure.DecoderConfig) {
		m.TagName = "json"
	}); err != nil {
		return nil, err
	}
	switch cfg.Target {
	case Stdout:
		return &Console{os.Stdout}, nil
	case Stderr:
		return &Console{os.Stderr}, nil
	default:
		return nil, fmt.Errorf("unknown target %q", cfg.Target)
	}

}

func init() {
	if err := appender.RegisterAppender("console", NewConsole); err != nil {
		panic(err)
	}
}
