package config

import "github.com/shanexu/logn/common"

type Config struct {
	Appenders map[string][]*common.Config `logn-config:"appenders"`
	Loggers   Loggers                     `logn-config:"loggers"`
}

var defaultConfig Config

func DefaultConfig() Config {
	return defaultConfig
}

type Loggers struct {
	Root   RootLogger `logn-config:"root"`
	Logger []Logger   `logn-config:"logger"`
}

type RootLogger struct {
	Level        string   `logn-config:"level"`
	AppenderRefs []string `logn-config:"appender_refs"`
}

type Logger struct {
	Name         string   `logn-config:"name" logn-validate:"required"`
	Level        string   `logn-config:"level"`
	AppenderRefs []string `logn-config:"appender_refs"`
}

func init() {
	defaultConfig = Config{
		Appenders: map[string][]*common.Config{
			"console": {common.MustNewConfigFrom(`
name: CONSOLE
target: stdout
encoder:
`)},
		},
		Loggers: Loggers{
			Root: RootLogger{
				AppenderRefs: []string{"CONSOLE"},
				Level:        "info",
			},
			Logger: nil,
		},
	}
}
