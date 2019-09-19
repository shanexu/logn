package logn_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/shanexu/logn"
)

const testConfigFile = "/tmp/test_logn.yaml"

func TestScan(t *testing.T) {
	if err := logn.InitWithConfigFile(testConfigFile); err != nil {
		t.Fatal(err)
	}
	log := logn.GetLogger("test")
	ticker := time.NewTicker(time.Second * 1)
	timer := time.NewTimer(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			log.Info("this is info")
			log.Debug("this is debug")
		case <-timer.C:
			return
		}
	}
}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	testConfigFileContent := `
appenders:
  console:
    - name: CONSOLE
      target: stdout
      encoder:
        console:
          time_encoder: epoch_millis
loggers:
  root:
    level: info
    appender_refs:
      - CONSOLE

scan: true
scan_period: 1s
`
	file, err := os.Create(testConfigFile)
	if err != nil {
		panic(err)
	}
	if _, err := file.WriteString(testConfigFileContent); err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(time.Second*1 + time.Millisecond*100)
		testConfigFileContent := `
appenders:
  console:
    - name: CONSOLE
      target: stdout
      encoder:
        console:
          time_encoder: epoch_millis
loggers:
  root:
    level: debug
    appender_refs:
      - CONSOLE

scan: true
scan_period: 1s
`
		file, err := os.Create(testConfigFile)
		if err != nil {
			panic(err)
		}
		if _, err := file.WriteString(testConfigFileContent); err != nil {
			panic(err)
		}
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	m.Run()
	fmt.Println("end")
}
