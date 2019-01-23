package logn

import (
	"log"
	"testing"
	"time"
)

func TestGetLogger(t *testing.T) {
	defer Sync()

	hello := GetLogger("hello")
	helloworld := GetLogger("helloworld")

	ti := time.NewTimer(time.Second * 10)

	out:
	for {
		select {
		case <- ti.C:
			break out
		default:
			hello.Info("hello")
			world := GetLogger("world")
			world.Debug("world")
			helloworld.Info("hello world")
			helloworld.Error("hell")
			log.Println("hello")
		}
	}
}