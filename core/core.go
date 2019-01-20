package core

import (
	"github.com/shanexu/logp/appender"
	"github.com/shanexu/logp/common"
	cfg "github.com/shanexu/logp/config"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Core struct {
	nameToLogger   sync.Map
	nameToAppender map[string]appender.Appender
	rootAppender   *appender.Appender
	rootLevel      zapcore.Level
}

//func (c *Core)GetLogger(name string) logn.Logger {
//	logger, ok := c.nameToLogger.Load(name)
//	if ok {
//
//	}
//}

func New(rawConfig *common.Config) (*Core, error) {
	config := cfg.DefaultConfig()
	err := rawConfig.Unpack(&config)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
