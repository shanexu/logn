package encoder

import "github.com/spf13/viper"

type Factory func(*viper.Viper) (Encoder, error)

var encoders = map[string]Factory{}
