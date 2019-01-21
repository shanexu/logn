package zap

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shanexu/logn/appender"
	"github.com/shanexu/logn/common"
	cfg "github.com/shanexu/logn/config"
	"github.com/shanexu/logn/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Core struct {
	nameToLogger     sync.Map
	nameToAppender   map[string]*appender.Appender
	rootAppenders    map[string]*appender.Appender
	rootLevel        zapcore.LevelEnabler
	rootLevelName    string
	rootAppenderRefs []string
}

func createLevel(level string) (zapcore.LevelEnabler, error) {
	var l zapcore.Level
	if err := l.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}
	return zap.NewAtomicLevelAt(l), nil
}

func (c *Core) putAppender(name string, a *appender.Appender) error {
	if name == "" {
		return errors.New("name should not be empty")
	}
	if a == nil {
		return errors.New("appender should not be nil")
	}
	if _, exist := c.nameToAppender[name]; exist {
		return fmt.Errorf("duplicated appender name %q", name)
	}
	c.nameToAppender[name] = a
	return nil
}

func (c *Core) getAppender(name string) (*appender.Appender, error) {
	a, exist := c.nameToAppender[name]
	if !exist {
		return nil, fmt.Errorf("not found appender %q", name)
	}
	return a, nil
}

func (c *Core) newLogger(loggerCfg cfg.Logger) (*ZapLogger, error) {
	name := loggerCfg.Name
	levelName := loggerCfg.Level
	afs := loggerCfg.AppenderRefs

	if levelName == "" {
		levelName = c.rootLevelName
	}

	if len(afs) == 0 {
		afs = c.rootAppenderRefs
	}

	level, err := createLevel(levelName)
	if err != nil {
		return nil, err
	}

	am := make(map[string]*appender.Appender)
	for a := range common.MakeStringSet(afs...) {
		var err error
		am[a], err = c.getAppender(a)
		if err != nil {
			return nil, err
		}
	}

	if len(am) == 0 {
		return nil, errors.New("empty appenders")
	}

	zcs := make([]zapcore.Core, 0)
	for _, a := range am {
		zcs = append(zcs, zapcore.NewCore(a.Encoder, a.Writer, level))
	}
	zt := zapcore.NewTee(zcs...)
	return NewZapLogger(zap.New(zt, zap.AddCaller(), zap.AddCallerSkip(1)).Named(name).Sugar()), nil
}

func (c *Core) newNamedLogger(name string) (z *ZapLogger) {
	zcs := make([]zapcore.Core, 0)
	for _, a := range c.rootAppenders {
		zcs = append(zcs, zapcore.NewCore(a.Encoder, a.Writer, c.rootLevel))
	}
	zt := zapcore.NewTee(zcs...)
	return NewZapLogger(zap.New(zt, zap.AddCaller(), zap.AddCallerSkip(1)).Named(name).Sugar())
}

func (c *Core) GetLogger(name string) core.Logger {
	logger, ok := c.nameToLogger.Load(name)
	if ok {
		return logger.(*ZapLogger)
	}
	zl := c.newNamedLogger(name)
	v, _ := c.nameToLogger.LoadOrStore(name, zl)
	return v.(*ZapLogger)
}

func New(rawConfig *common.Config) (core.Core, error) {
	config := cfg.DefaultConfig()
	err := rawConfig.Unpack(&config)
	if err != nil {
		return nil, err
	}

	co := Core{
		nameToLogger:   sync.Map{},
		nameToAppender: map[string]*appender.Appender{},
		rootAppenders:  map[string]*appender.Appender{},
	}

	for appenderType, appenderConfigs := range config.Appenders {
		for _, appenderConfig := range appenderConfigs {
			a, err := appender.CreateAppender(appenderType, appenderConfig)
			if err != nil {
				return nil, err
			}
			name, err := appenderConfig.Name()
			if err != nil {
				return nil, err
			}
			if err := co.putAppender(name, a); err != nil {
				return nil, err
			}
		}
	}

	// rootLevel
	rootLevel, err := createLevel(config.Loggers.Root.Level)
	if err != nil {
		return nil, err
	}
	co.rootLevel = rootLevel
	co.rootLevelName = config.Loggers.Root.Level

	// rootAppenders
	rootAppenderRefSet := common.MakeStringSet(config.Loggers.Root.AppenderRefs...)
	for appenderRef := range rootAppenderRefSet {
		a, err := co.getAppender(appenderRef)
		if err != nil {
			return nil, err
		}
		co.rootAppenders[appenderRef] = a
	}
	co.rootAppenderRefs = rootAppenderRefSet.ToSlice()

	// loggers
	for _, lc := range config.Loggers.Logger {
		l, err := co.newLogger(lc)
		if err != nil {
			return nil, err
		}
		if _, loaded := co.nameToLogger.LoadOrStore(lc.Name, l); loaded {
			return nil, fmt.Errorf("duplicated logger %q", lc.Name)
		}
	}

	return &co, nil
}

func (c *Core) Sync() {
	c.nameToLogger.Range(func(_, value interface{}) bool {
		value.(*ZapLogger).Sync()
		return false
	})
	for _, a := range c.rootAppenders {
		a.Writer.Sync()
	}
}

func init()  {
	core.RegisterType("zap", New)
	core.RegisterType("default", New)
}