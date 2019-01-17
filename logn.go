package logn

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

var (
	_nameToLogger = sync.Map{}
)

func init() {
	v := viper.New()
	configFile := os.Getenv("LOGN_CONFIG")
	if configFile != "" {
		ext := filepath.Ext(configFile)
		if ext != "" {
			ext = ext[1:]
		}
		if ext != "" {
			v.SetConfigType(ext)
		}
		v.SetConfigFile(configFile)
	}
	if configFile == "" {
		v.AddConfigPath(".")
		v.SetConfigName("logn")
	}
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("logn using config file:", v.ConfigFileUsed())
}

func GetLogger(name string) Logger {
	l, found := _nameToLogger.Load(name)
	if found {
		return l.(Logger)
	}
	l, _ = _nameToLogger.LoadOrStore(name, newLogger(name))
	return l.(Logger)
}

func Root() Logger {
	return &ZapLogger{}
}

func newLogger(name string) Logger {
	return &ZapLogger{}
}
