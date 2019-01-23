package logn

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	defer Sync()

	hello := GetLogger("hello")
	hello.Info("hello")
	world := GetLogger("world")
	world.Debug("world")

	helloworld := GetLogger("helloworld")
	helloworld.Info("hello world")
	helloworld.Error("hell")
}