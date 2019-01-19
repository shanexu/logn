package encoder

import (
	"fmt"
	"github.com/spf13/viper"
)

type Factory func(*viper.Viper) (Encoder, error)

var encoders = map[string]Factory{}

func RegisterType(name string, gen Factory) {
	if _, exists := encoders[name]; exists {
		panic(fmt.Sprintf("encoder %q already registered", name))
	}
}

func CreateEncoder(cfg viper.Viper) (Encoder, error) {
	return nil, nil
}