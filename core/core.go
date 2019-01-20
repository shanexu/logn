package core

import (
	"github.com/shanexu/logp/appender"
	"github.com/shanexu/logp/common"
	cfg "github.com/shanexu/logp/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Core struct {
	nameToLogger   sync.Map
	nameToAppender map[string]appender.Appender
	rootAppender   *appender.Appender
	rootLevel      zapcore.LevelEnabler
}

func (c *Core) newNamedLogger(name string) *ZapLogger {
	zc := zapcore.NewCore(c.rootAppender.Encoder, c.rootAppender.Writer, c.rootLevel)
	l := zap.New(zc).Named(name).Sugar()
	return NewZapLogger(l)
}

func (c *Core) GetLogger(name string) *ZapLogger {
	logger, ok := c.nameToLogger.Load(name)
	if ok {
		return logger.(*ZapLogger)
	}
	zl := c.newNamedLogger(name)
	v, _ := c.nameToLogger.LoadOrStore(name, zl)
	return v.(*ZapLogger)
}

func New(rawConfig *common.Config) (*Core, error) {
	config := cfg.DefaultConfig()
	err := rawConfig.Unpack(&config)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
