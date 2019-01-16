package file

import (
	"github.com/shanexu/logp/appender"
	"github.com/spf13/viper"
	"os"
)

type File struct {
	*os.File
}

type Config struct {
	FileName string `json:"file_name"`
}

var (
	defaultConfig = Config{}
)

func DefaultConfig() Config {
	return defaultConfig
}

func NewFile(v viper.Viper) (appender.Appender, error) {
	cfg := DefaultConfig()
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(cfg.FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return &File{f}, nil
}

func init() {
	if err := appender.RegisterAppender("file", NewFile); err != nil {
		panic(err)
	}
}
