package logn

import (
	"fmt"
	"github.com/shanexu/logn/common"
	"github.com/shanexu/logn/core"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	_ "github.com/shanexu/logn/includes"
)

var logncore core.Core

func init() {
	configFile := os.Getenv("LOGN_CONFIG")
	debug := os.Getenv("LOGN_DEBUG")

	if configFile == "" {
		matches1, _ := filepath.Glob("logn.yaml")
		matches2, _ := filepath.Glob("logn.yml")
		matches := append(matches1, matches2...)
		switch len(matches) {
		case 0:
			if debug == "true" {
				fmt.Println("no config file found using default config")
			}
		case 1:
			configFile = matches[0]
		default:
			panic(fmt.Errorf("multi config files found %v", matches))
		}
	}

	var err error
	var rawConfig *common.Config

	if configFile != "" {
		// load ConfigFile
		configFile, err = filepath.Abs(configFile)
		if err != nil {
			panic(err)
		}

		if debug == "true" {
			fmt.Println("logn using config file:", configFile)
			bs, err := ioutil.ReadFile(configFile)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bs))
		}

		rawConfig, err = common.LoadFile(configFile)
	} else {
		// load default config
		rawConfig, err = common.NewConfigFrom(`
appenders:
  console:
    - name: CONSOLE
      target: stdout
      encoder:
        console:
          time_encoder: ISO8601
loggers:
  root:
    level: info
    appender_refs:
      - CONSOLE
`)
	}

	if err != nil {
		panic(err)
	}

	co, err := core.CreateCore(rawConfig)

	if err != nil {
		panic(err)
	}
	logncore = co

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		Sync()
	}()
}

func Sync() {
	logncore.Sync()
}

func GetLogger(name string) core.Logger {
	return logncore.GetLogger(name)
}
