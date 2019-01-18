package logn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/shanexu/logp/appender"
	"github.com/shanexu/logp/configuration"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

var (
	_nameToLogger   = sync.Map{}
	_nameToAppender = make(map[string]appender.Appender)
	_rootLogger     Logger
	_rootAppender   appender.Appender
	_rootLevel      zapcore.LevelEnabler
)

var (
	errNoAppender = errors.New("no appender")
)

func newRootLogger() Logger {
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		_rootAppender,
		_rootLevel)
	return &ZapLogger{zap.New(core).Sugar()}
}

func newLevelEnabler(l string) (zapcore.LevelEnabler, error) {
	level := new(zapcore.Level)
	if err := level.UnmarshalText([]byte(l)); err != nil {
		return nil, err
	}
	return zap.NewAtomicLevelAt(*level), nil
}

func newAppender(aps []string) (appender.Appender, error) {
	appendersMap := make(map[string]appender.Appender)
	for _, aname := range aps {
		a, ok := _nameToAppender[aname]
		if !ok {
			return nil, fmt.Errorf("append %q not found", aname)
		}
		appendersMap[aname] = a
	}
	if len(appendersMap) == 0 {
		return nil, errNoAppender
	}
	appenders := make([]zapcore.WriteSyncer, 0)
	for _, a := range appendersMap {
		appenders = append(appenders, a)
	}
	return zapcore.NewMultiWriteSyncer(appenders...), nil
}

func newLogger(cfg configuration.Logger) (Logger, error) {
	a := _rootAppender
	if len(cfg.AppenderRefs) > 0 {
		ap, err := newAppender(cfg.AppenderRefs)
		if err != nil {
			return nil, err
		}
		a = ap
	}
	l := _rootLevel
	if cfg.Level != "" {
		le, err := newLevelEnabler(cfg.Level)
		if err != nil {
			return nil, err
		}
		l = le
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), a, l)
	return zap.New(core).With(zap.String("logger", cfg.Name)).Sugar(), nil
}

func init() {
	v := viper.New()
	configFile := os.Getenv("LOGN_CONFIG")
	debug := os.Getenv("LOG_CONFIG_DEBUG")
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
	cfg := configuration.Configuration{}
	err = v.Unmarshal(&cfg, func(m *mapstructure.DecoderConfig) {
		m.TagName = "json"
	})
	if err != nil {
		panic(err)
	}

	if debug == "true" {
		buf := bytes.NewBuffer(nil)
		bs, _ := json.Marshal(cfg)
		_ = json.Indent(buf, bs, "", "  ")
		fmt.Println(buf.String())
	}

	for atype, appenders := range cfg.Appenders {
		for _, config := range appenders {
			vv := viper.New()
			if err := vv.MergeConfigMap(config); err != nil {
				panic(err)
			}
			a, err := appender.NewAppender(atype, vv)
			if err != nil {
				panic(err)
			}
			if _, found := _nameToAppender[config.Name()]; found {
				panic(fmt.Errorf("duplicated appender name %q", config.Name()))
			}
			_nameToAppender[config.Name()] = a
		}
	}

	cfg.Loggers.Root.Name = "__root__"
	_rootAppender, err = newAppender(cfg.Loggers.Root.AppenderRefs)
	if err != nil {
		panic(err)
	}
	_rootLevel, err = newLevelEnabler(cfg.Loggers.Root.Level)
	if err != nil {
		panic(err)
	}
	_rootLogger = newRootLogger()

	for _, c := range cfg.Loggers.Logger {
		l, err := newLogger(c)
		if err != nil {
			panic(err)
		}
		if _, loaded := _nameToLogger.LoadOrStore(c.Name, l); loaded {
			panic(fmt.Errorf("duplicated logger %q", c.Name))
		}
	}

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		Sync()
	}()
}

func Sync() {
	for _, w := range _nameToAppender {
		w.Sync()
	}
}

func GetLogger(name string) Logger {
	l, found := _nameToLogger.Load(name)
	if found {
		return l.(Logger)
	}
	l, _ = _nameToLogger.LoadOrStore(name, newNamedLogger(name))
	return l.(Logger)
}

func Root() Logger {
	return &ZapLogger{}
}

func newNamedLogger(name string) Logger {
	return &ZapLogger{_rootLogger.(*ZapLogger).sugar.With("logger", name)}
}
