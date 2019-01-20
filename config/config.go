package config

import "github.com/shanexu/logp/common"

type Config struct {
	Appenders map[string][]*common.Config `config:"appenders"`
	Loggers   Loggers                     `config:"loggers"`
}

var defaultConfig Config

func DefaultConfig() Config {
	return defaultConfig
}

type Loggers struct {
	Root   RootLogger `config:"root"`
	Logger []Logger   `config:"logger"`
}

type RootLogger struct {
	Level        string   `config:"level"`
	AppenderRefs []string `config:"appender_refs"`
}

type Logger struct {
	Name         string   `config:"name"`
	AppenderRefs []string `config:"appender_refs"`
}

func init() {
	defaultConfig = Config{
		Appenders: map[string][]*common.Config{
			"console": {common.MustNewConfigFrom(`
name: CONSOLE
target: stdout
encoder:
  json:
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
