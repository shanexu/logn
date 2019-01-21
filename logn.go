package logn

import (
	"errors"
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

var logncore *core.Core

func init() {
	configFile := os.Getenv("LOGN_CONFIG")
	debug := os.Getenv("LOGN_DEBUG")

	if configFile == "" {
		matches1, _ := filepath.Glob("logn.yaml")
		matches2, _ := filepath.Glob("logn.yml")
		matches := append(matches1, matches2...)
		switch len(matches) {
		case 0:
			panic(errors.New("no config file found"))
		case 1:
			configFile = matches[0]
		default:
			panic(fmt.Errorf("multi config files found %v", matches))
		}
	}

	var err error
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

	rawConfig, err := common.LoadFile(configFile)
	if err != nil {
		panic(err)
	}

	co, err := core.New(rawConfig)
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

func GetLogger(name string) Logger {
	return logncore.GetLogger(name)
}
