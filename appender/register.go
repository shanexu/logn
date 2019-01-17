package appender

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	_appenderNameToConstructor = make(map[string]func(*viper.Viper) (Appender, error), 0)
	_appenderMutex sync.RWMutex
)

var (
	errNoAppenderNameSpecified = errors.New("no appender name specified")
)

func RegisterAppender(name string, constructor func(*viper.Viper) (Appender, error)) error {
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

func NewAppender(atype string, v *viper.Viper) (Appender, error){
	_appenderMutex.RLock()
	defer _appenderMutex.RUnlock()
	c, ok := _appenderNameToConstructor[atype]
	if !ok {
		return nil, fmt.Errorf("appender type %q not found", atype)
	}
	return c(v)
}