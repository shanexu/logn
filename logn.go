package logn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/shanexu/logp/appender/writer"
	"github.com/shanexu/logp/common"
	"github.com/shanexu/logp/config"
	"github.com/shanexu/logp/configuration"
	"github.com/shanexu/logp/core"
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
	_nameToAppender = make(map[string]writer.Writer)
	_rootLogger     Logger
	_rootAppender   writer.Writer
	_rootLevel      zapcore.LevelEnabler
)

var (
	errNoAppender = errors.New("no appender")
)

func newRootLogger() Logger {
	zc := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		_rootAppender,
		_rootLevel)
	return core.NewZapLogger(zap.New(zc).Sugar())
}

func newLevelEnabler(l string) (zapcore.LevelEnabler, error) {
	level := new(zapcore.Level)
	if err := level.UnmarshalText([]byte(l)); err != nil {
		return nil, err
	}
	return zap.NewAtomicLevelAt(*level), nil
}

func newAppender(aps []string) (writer.Writer, error) {
	appendersMap := make(map[string]writer.Writer)
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

	if debug == "true" {
		fmt.Println("logn using config file:", v.ConfigFileUsed())
	}

	cfg2 := configuration.Configuration{}
	err = v.Unmarshal(&cfg2, func(m *mapstructure.DecoderConfig) {
		m.TagName = "json"
	})
	if err != nil {
		panic(err)
	}

	cfg, err := common.LoadFile(v.ConfigFileUsed())
	if err != nil {
		panic(err)
	}

	if debug == "true" {
		buf := bytes.NewBuffer(nil)
		bs, _ := json.Marshal(cfg2)
		_ = json.Indent(buf, bs, "", "  ")
		fmt.Println(buf.String())
	}

	config := config.Config{}
	if err := cfg.Unpack(&config); err != nil {
		panic(err)
	}
	for atype, appenders := range config.Appenders {
		for _, v := range appenders {
			name, _ := v.String("name", -1)
			appender, err := writer.NewWriter(atype, v)
			if err != nil {
				panic(err)
			}
			if _, exist := _nameToAppender[name]; exist {
				panic(fmt.Errorf("dublicated appender %q", name))
			}
			_nameToAppender[name] = appender
		}
	}

	cfg2.Loggers.Root.Name = "__root__"
	_rootAppender, err = newAppender(cfg2.Loggers.Root.AppenderRefs)
	if err != nil {
		panic(err)
	}
	_rootLevel, err = newLevelEnabler(cfg2.Loggers.Root.Level)
	if err != nil {
		panic(err)
	}
	_rootLogger = newRootLogger()

	for _, c := range cfg2.Loggers.Logger {
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
	_nameToLogger.Range(func(_, value interface{}) bool {
		_ = value.(Logger).Sync()
		return true
	})
	for _, w := range _nameToAppender {
		_ = w.Sync()
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
	return &core.ZapLogger{}
}

func newNamedLogger(name string) Logger {
	l, _ := newLogger(configuration.Logger{Name: name})
	return l
}
