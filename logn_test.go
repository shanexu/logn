package logn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func TestGetLogger(t *testing.T) {
	if os.Getenv("LONG_TEST") != "on" {
		t.Skip("skip long test")
	}
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

	rootLogger := GetLogger()

out:
	for {
		select {
		case <-ti.C:
			break out
		case <-tc.C:
			hello.Info("hello")
			world.Debug("world")
			helloworld.Info("hello world")
			helloworld.Error(hell)
			log.Println("hello")
			rootLogger.Info("root logger")
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

func TestGlobalLog(t *testing.T) {
	Info("hello")
	Error("hell")
}

func TestRedirectStdLog(t *testing.T) {
	log.Println("hello")
}

func TestInitWithConfigContent(t *testing.T) {
	const newConfig = `appenders:
  console:
    - name: CONSOLE
      target: stdout
      encoder:
        console:
loggers:
  root:
    level: info
    appender_refs:
      - CONSOLE`
	l := GetLogger("l")
	Info("hello")
	l.Info("hello")
	err := InitWithConfigContent(newConfig)
	assert.Nil(t, err)
	Info("world")
	l.Info("world")
	err = InitWithConfigContent(newConfig)
	assert.NotNil(t, err)
}

func TestInitWithConfigFile(t *testing.T) {
	l := GetLogger("l")
	Info("hello")
	l.Info("hello")
	err := InitWithConfigFile("test/.logn.yml")
	Info("hello")
	l.Info("hello")
	err = InitWithConfigFile("test/.logn.yml")
	assert.NotNil(t, err)
}

func TestInitWithConfigFile2(t *testing.T) {
	err := InitWithConfigFile("test/.logn_not_exist.yml")
	assert.NotNil(t, err)
}
