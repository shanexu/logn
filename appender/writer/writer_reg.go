package writer

import (
	"fmt"
	"github.com/shanexu/logp/common"
)

type Factory func(config *common.Config) (Writer, error)

var (
	writers = map[string]Factory{}
)

func RegisterType(name string, f Factory) {
	if writers[name] != nil {
		panic(fmt.Errorf("writer type  '%v' exists already", name))
	}
	writers[name] = f
}

func NewWriter(name string, config *common.Config) (Writer, error) {
	factory := writers[name]
	if factory == nil {
		return nil, fmt.Errorf("writer type %v undefined", name)
	}
	return factory(config)
}
