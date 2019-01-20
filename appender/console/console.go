package console

import (
	"fmt"
	"github.com/shanexu/logp/appender/writer"
	"github.com/shanexu/logp/common"
	"os"
)

type Console struct {
	*os.File
}

type Config struct {
	Target `config:"target" validate:"required,oneof=stderr stdout"`
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

func NewConsole(v *common.Config) (writer.Writer, error) {
	cfg := DefaultConfig()
	if err := v.Unpack(&cfg); err != nil {
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
	writer.RegisterType("console", NewConsole)
}
