package appender

import (
	"errors"
	"fmt"
	"github.com/shanexu/logp/common"
	"github.com/spf13/viper"
	"sync"
)

type WriterFactory func(config *common.Config) (Writer, error)

var (
	_appenderNameToConstructor = make(map[string]func(*viper.Viper) (Writer, error))
	_appenderMutex             sync.RWMutex
	writers                    = map[string]WriterFactory{}
)

var (
	errNoAppenderNameSpecified = errors.New("no appender name specified")
)

func RegisterType(name string, f WriterFactory) {
	if writers[name] != nil {
		panic(fmt.Errorf("writer type  '%v' exists already", name))
	}
	writers[name] = f
}

func CreateWriter(name string, config *common.Config) (Writer, error) {
	factory := writers[name]
	if factory == nil {
		return nil, fmt.Errorf("writer type %v undefined", name)
	}
	return factory(config)
}

func RegisterAppender(name string, constructor func(*viper.Viper) (Writer, error)) error {
	_appenderMutex.Lock()
	defer _appenderMutex.Unlock()
	if name == "" {
		return errNoAppenderNameSpecified
	}
	if _, ok := _appenderNameToConstructor[name]; ok {
		return fmt.Errorf("appender already registered for name %q", name)
	}
	_appenderNameToConstructor[name] = constructor
	return nil
}

func NewAppender(atype string, v *viper.Viper) (Writer, error) {
	_appenderMutex.RLock()
	defer _appenderMutex.RUnlock()
	c, ok := _appenderNameToConstructor[atype]
	if !ok {
		return nil, fmt.Errorf("appender type %q not found", atype)
	}
	return c(v)
}
