package config

import "github.com/shanexu/logp/common"

type Config struct {
	Appenders map[string][]*common.Config `config:"appenders"`
	Loggers   Loggers                     `config:"loggers"`
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
