package config

import (
	"github.com/shanexu/logn/common"
)

type Config struct {
	Appenders map[string][]*common.Config `logn-config:"appenders"`
	Loggers   Loggers                     `logn-config:"loggers"`
}

type ScanConfig struct {
	Scan       bool   `logn-config:"scan"`
	ScanPeriod string `logn-config:"scan_period"`
}

type Loggers struct {
	Root   RootLogger `logn-config:"root"`
	Logger []Logger   `logn-config:"logger"`
}

type RootLogger struct {
	Level        string   `logn-config:"level"`
	CallerSkip   *int     `logn-config:"caller_skip"`
	AppenderRefs []string `logn-config:"appender_refs"`
}

type Logger struct {
	Name         string   `logn-config:"name" logn-validate:"required"`
	Level        string   `logn-config:"level"`
	CallerSkip   *int     `logn-config:"caller_skip"`
	AppenderRefs []string `logn-config:"appender_refs"`
}
