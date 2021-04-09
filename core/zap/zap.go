package zap

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/shanexu/logn/appender"
	"github.com/shanexu/logn/common"
	cfg "github.com/shanexu/logn/config"
	"github.com/shanexu/logn/core"
)

type Core struct {
	locker           sync.RWMutex
	nameToLogger     sync.Map
	nameToAppender   map[string]*appender.Appender
	rootAppenders    map[string]*appender.Appender
	rootLevel        zapcore.LevelEnabler
	rootLevelName    string
	rootCallerSkip   *int
	rootAppenderRefs []string
	rootLogger       *zap.SugaredLogger
	globalLogger     *zap.SugaredLogger
}

var StackTraceLevelEnabler = zap.NewAtomicLevelAt(zapcore.ErrorLevel)

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

func (c *Core) getAppenders(names []string) (map[string]*appender.Appender, error) {
	m := make(map[string]*appender.Appender)
	for _, name := range names {
		a, err := c.getAppender(name)
		if err != nil {
			return nil, err
		}
		m[name] = a
	}
	return m, nil
}

func newZapCore(level zapcore.LevelEnabler, appenders map[string]*appender.Appender) zapcore.Core {
	zcs := make([]zapcore.Core, 0)
	for _, a := range appenders {
		zcs = append(zcs, zapcore.NewCore(a.Encoder, a.Writer, level))
	}
	return zapcore.NewTee(zcs...)
}

func newLogger(name string, level zapcore.LevelEnabler, callerSkip *int, appenders map[string]*appender.Appender) *zap.SugaredLogger {
	zc := newZapCore(level, appenders)
	cs := 0
	if callerSkip != nil {
		cs = *callerSkip
	}
	logger := zap.New(zc, zap.AddCaller(), zap.AddStacktrace(StackTraceLevelEnabler), zap.AddCallerSkip(cs))
	if name != "" {
		logger = logger.Named(name)
	}
	return logger.Sugar()
}

func (c *Core) newLoggerFromCfg(loggerCfg cfg.Logger) (core.Logger, error) {
	name := loggerCfg.Name
	levelName := loggerCfg.Level
	afs := loggerCfg.AppenderRefs
	callerSkip := loggerCfg.CallerSkip

	if levelName == "" {
		levelName = c.rootLevelName
	}

	if len(afs) == 0 {
		afs = c.rootAppenderRefs
	}

	if callerSkip == nil {
		callerSkip = loggerCfg.CallerSkip
	}

	level, err := createLevel(levelName)
	if err != nil {
		return nil, err
	}

	am, err := c.getAppenders(afs)
	if err != nil {
		return nil, err
	}

	if len(am) == 0 {
		return nil, errors.New("empty appenders")
	}

	return newLogger(name, level, callerSkip, am), nil
}

func (c *Core) newNamedLogger(name string) core.Logger {
	return newLogger(name, c.rootLevel, c.rootCallerSkip, c.rootAppenders)
}

func (c *Core) getLogger(name string, lock bool) *zap.SugaredLogger {
	if lock {
		c.locker.RLock()
		defer c.locker.RUnlock()
	}
	if len(name) == 0 {
		return c.rootLogger
	}
	logger, ok := c.nameToLogger.Load(name)
	if ok {
		return logger.(*zap.SugaredLogger)
	}
	zl := c.newNamedLogger(name)
	v, _ := c.nameToLogger.LoadOrStore(name, zl)
	return v.(*zap.SugaredLogger)
}

func (c *Core) GetLogger(name ...string) core.Logger {
	if len(name) == 0 {
		return c.getLogger("", true)
	}
	return c.getLogger(name[0], true)
}

func (c *Core) Update(rawConfig *common.Config) error {
	nc, err := newCore(rawConfig)
	if err != nil {
		return err
	}
	c.locker.Lock()
	defer c.locker.Unlock()
	c.Sync()
	c.nameToAppender = nc.nameToAppender
	c.rootAppenders = nc.rootAppenders
	c.rootLevel = nc.rootLevel
	c.rootLevelName = nc.rootLevelName
	c.rootCallerSkip = nc.rootCallerSkip
	c.rootAppenderRefs = nc.rootAppenderRefs
	*c.rootLogger = *nc.rootLogger
	c.rootAppenders = nc.rootAppenders
	c.globalLogger = nc.globalLogger
	c.nameToLogger.Range(func(key, value interface{}) bool {
		name := key.(string)
		*value.(*zap.SugaredLogger) = *nc.getLogger(name, false)
		return true
	})
	nc.nameToLogger.Range(func(key, value interface{}) bool {
		if _, found := c.nameToLogger.Load(key); found {
			return true
		}
		c.nameToLogger.Store(key, value)
		return true
	})
	c.redirectStdLog()
	return nil
}

func newCore(rawConfig *common.Config) (*Core, error) {
	config := cfg.Config{}
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

	co.rootCallerSkip = config.Loggers.Root.CallerSkip

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

	// rootLogger
	co.rootLogger = newLogger("", co.rootLevel, co.rootCallerSkip, co.rootAppenders)

	// loggers
	for _, lc := range config.Loggers.Logger {
		l, err := co.newLoggerFromCfg(lc)
		if err != nil {
			return nil, err
		}
		if _, loaded := co.nameToLogger.LoadOrStore(lc.Name, l); loaded {
			return nil, fmt.Errorf("duplicated logger %q", lc.Name)
		}
	}

	co.globalLogger = co.rootLogger.With().Desugar().WithOptions(zap.AddCallerSkip(2)).Sugar()

	return &co, nil
}

func New(rawConfig *common.Config) (core.Core, error) {
	return newCore(rawConfig)
}

func (c *Core) RedirectStdLog() {
	zap.RedirectStdLog(c.getLogger("stdlog", true).Desugar())
}

func (c *Core) redirectStdLog() {
	zap.RedirectStdLog(c.getLogger("stdlog", false).Desugar())
}

func (c *Core) Sync() error {
	c.nameToLogger.Range(func(_, value interface{}) bool {
		value.(core.Logger).Sync()
		return false
	})
	for _, a := range c.rootAppenders {
		a.Writer.Sync()
	}
	return nil
}

func (c *Core) Global() core.Logger {
	return c.globalLogger
}

func init() {
	core.RegisterType("zap", New)
	core.RegisterType("default", New)
}
