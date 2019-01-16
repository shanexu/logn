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
	FileName string      `json:"file_name"`
	Perm     os.FileMode `json:"perm"`
}

var (
	defaultConfig = Config{
		Perm: os.ModePerm,
	}
)

func DefaultConfig() Config {
	return defaultConfig
}

func NewFile(v viper.Viper) (appender.Appender, error) {
	return &File{}, nil
}

func init() {
	if err := appender.RegisterAppender("file", NewFile); err != nil {
		panic(err)
	}
}
