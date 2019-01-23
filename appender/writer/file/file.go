package file

import (
	"github.com/shanexu/logn/appender/writer"
	"github.com/shanexu/logn/common"
	"os"
)

type File struct {
	*os.File
}

type Config struct {
	FileName string `logn-config:"file_name" logn-validate:"required"`
}

var (
	defaultConfig = Config{}
)

func DefaultConfig() Config {
	return defaultConfig
}

func NewFile(v *common.Config) (writer.Writer, error) {
	cfg := DefaultConfig()
	if err := v.Unpack(&cfg); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(cfg.FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &File{f}, nil
}

func init() {
	writer.RegisterType("file", NewFile)
}
