package logn

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestGetLogger(t *testing.T) {
	defer Sync()

	hello := GetLogger("hello")
	helloworld := GetLogger("helloworld")
	world := GetLogger("world")

	ti := time.NewTimer(time.Second * 10)
	tc := time.NewTicker(time.Millisecond * 10)

	hell := ""
	for i := 0; i < 100; i++ {
		hell += fmt.Sprintf("hell%d", i)
	}

	out:
	for {
		select {
		case <- ti.C:
			break out
		case <- tc.C:
			hello.Info("hello")
			world.Debug("world")
			helloworld.Info("hello world")
			helloworld.Error(hell)
			log.Println("hello")
		}
	}
}

func TestGetLogger2(t *testing.T) {
	helloworld := GetLogger("helloworld")
	hell := ""
	for i := 0; i < 10000; i++ {
		hell += fmt.Sprintf("hell%d", i)
	}
	helloworld.Error(hell)
}