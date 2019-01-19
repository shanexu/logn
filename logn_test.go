package logn

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestGetLogger(t *testing.T) {
	hello := GetLogger("hello")
	hello.Info("hello")
	world := GetLogger("world")
	world.Debug("world")

	helloworld := GetLogger("helloworld")
	helloworld.Info("hello world")
	Sync()
}

func TestReadConfig(t *testing.T) {
	v := viper.New()
	v.SetConfigType("yml")
	err := v.ReadConfig(bytes.NewReader([]byte(`
- a:
   b: 1
`)))
	if err != nil {
		fmt.Println(err)
	}
}