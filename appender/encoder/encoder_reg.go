package encoder

import (
	"fmt"
	"github.com/shanexu/logp/common"
)

type Factory func(*common.Config) (Encoder, error)

type Config struct {
	Namespace common.ConfigNamespace `config:",inline"`
}

var encoders = map[string]Factory{}

func RegisterType(name string, gen Factory) {
	if _, exists := encoders[name]; exists {
		panic(fmt.Sprintf("encoder %q already registered", name))
	}
	encoders[name] = gen
}

func CreateEncoder(cfg Config) (Encoder, error) {
	// default to json encoder
	encoder := "json"
	if name := cfg.Namespace.Name(); name != "" {
		encoder = name
	}

	factory := encoders[encoder]
	if factory == nil {
		return nil, fmt.Errorf("'%v' encoder is not available", encoder)
	}
	return factory(cfg.Namespace.Config())
}
