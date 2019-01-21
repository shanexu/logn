package core

import (
	"fmt"
	"github.com/shanexu/logn/common"
)

type Factory func(config *common.Config) (Core, error)

var (
	cores = map[string]Factory{}
)

func RegisterType(name string, f Factory) {
	if cores[name] != nil {
		panic(fmt.Errorf("core type %q exists already", name))
	}
	cores[name] = f
}

func CreateCore(config *common.Config) (Core, error) {
	 backend, _ := config.String("backend", -1)
	 if backend == "" {
	 	backend = "default"
	 }
	 f, _ := cores[backend]
	 if f == nil {
		 return nil, fmt.Errorf("no core backend %q available", backend)
	 }
	 return f(config)
}