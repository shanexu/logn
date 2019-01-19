package file

import (
	"github.com/shanexu/logp/appender"
	"github.com/shanexu/logp/common"
	"os"
)

type File struct {
	*os.File
}

type Config struct {
	FileName string `config:"file_name" validate:"required"`
}

var (
	defaultConfig = Config{}
)

func DefaultConfig() Config {
	return defaultConfig
}

func NewFile(v *common.Config) (appender.Writer, error) {
	cfg := DefaultConfig()
	if err := v.Unpack(&cfg); err != nil {
		return nil, err
	}
	if err := common.Validate().Struct(cfg); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(cfg.FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &File{f}, nil
}

func init() {
	appender.RegisterType("file", NewFile)
}
