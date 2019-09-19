package logn

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"

	"github.com/shanexu/logn/common"
	"github.com/shanexu/logn/config"
	"github.com/shanexu/logn/core"

	_ "github.com/shanexu/logn/includes"
)

const DefaultConfig = `appenders:
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

scan: false
scan_period: 1m
`

var (
	logncore       core.Core
	configFile     string
	initLocker     sync.Mutex
	explicitInited = false
	debug          bool
)

func ConfigWithRawConfig(rawConfig *common.Config) (core.Core, error) {
	co, err := core.CreateCore(rawConfig)

	if err != nil {
		return nil, err
	}
	return co, nil
}

func resolveConfigFileFromEnv() (string, error) {
	f := os.Getenv("LOGN_CONFIG")
	if f == "" {
		return "", errors.New("environment variable 'LOGN_CONFIG' is not set")
	}
	return f, nil
}

func resolveConfigFileFromWorkDir() (string, error) {
	matches1, _ := filepath.Glob("logn.yaml")
	matches2, _ := filepath.Glob("logn.yml")
	matches := append(matches1, matches2...)
	switch len(matches) {
	case 0:
		return "", errors.New("no config file found in work dir")
	case 1:
		return matches[0], nil
	default:
		panic(fmt.Errorf("multiple config files found %v", matches))
	}
}

func InitWithConfigFile(path string) error {
	initLocker.Lock()
	defer initLocker.Unlock()

	if explicitInited {
		return errors.New("logn is explicit inited")
	}

	if path == "" {
		return errors.New("config file path is empty")
	}

	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		return err
	}

	if debug {
		fmt.Println("logn using config file: ", path)
		bs, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bs))
	}

	rawConfig, configFileHash, err := common.LoadFile(path)
	if err != nil {
		return err
	}

	err = logncore.Update(rawConfig)
	if err != nil {
		return err
	}

	configFile = path
	explicitInited = true

	watchConfigFile(configFile, configFileHash, rawConfig)

	return nil
}
func InitWithConfigContent(content string) error {
	initLocker.Lock()
	defer initLocker.Unlock()

	if explicitInited {
		return errors.New("logn is explicit inited")
	}

	if debug {
		fmt.Println("logn InitWithConfigContent:\n" + content)
	}

	rawConfig, err := common.NewConfigFrom(content)
	if err != nil {
		return err
	}

	err = logncore.Update(rawConfig)
	if err != nil {
		return err
	}

	explicitInited = true

	return nil
}

func init() {
	initLocker.Lock()
	defer initLocker.Unlock()

	debug = os.Getenv("LOGN_DEBUG") == "true"

	if configFile == "" {
		cf, err := resolveConfigFileFromEnv()
		if err == nil {
			configFile = cf
		}
	}

	if configFile == "" {
		cf, err := resolveConfigFileFromWorkDir()
		if err == nil {
			configFile = cf
		}
	}

	var err error
	var rawConfig *common.Config
	var configFileHash [md5.Size]byte

	if configFile != "" {
		// load ConfigFile
		configFile, err = filepath.Abs(configFile)
		if err != nil {
			panic(err)
		}

		if debug {
			fmt.Println("logn using config file: ", configFile)
			bs, err := ioutil.ReadFile(configFile)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bs))
		}

		rawConfig, configFileHash, err = common.LoadFile(configFile)
	} else {
		if debug {
			fmt.Print("logn using default config:\n" + DefaultConfig)
		}
		rawConfig, err = common.NewConfigFrom(DefaultConfig)
	}

	if err != nil {
		panic(err)
	}

	co, err := ConfigWithRawConfig(rawConfig)

	if err != nil {
		panic(err)
	}

	logncore = co
	logncore.RedirectStdLog()

	if configFile != "" {
		explicitInited = true
	}

	if explicitInited {
		watchConfigFile(configFile, configFileHash, rawConfig)
	}

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		Sync()
	}()
}

func watchConfigFile(configFile string, configFileHash [md5.Size]byte, rawConfig *common.Config) {
	scanConfig := config.ScanConfig{
		Scan:       false,
		ScanPeriod: "1m",
	}
	if err := rawConfig.Unpack(&scanConfig); err != nil {
		panic(err)
	}
	if scanConfig.Scan {
		scanPeriod, err := time.ParseDuration(scanConfig.ScanPeriod)
		if err != nil {
			panic(err)
		}
		defer func() {
			for {
				t := time.NewTimer(scanPeriod)
				<-t.C
				rawConfig, hash, err := common.LoadFile(configFile)
				if err != nil {
					continue
				}
				if configFileHash != hash {
					configFileHash = hash
					if err := logncore.Update(rawConfig); err == nil {
						scanConfig = config.ScanConfig{
							Scan:       false,
							ScanPeriod: "1m",
						}
						if err := rawConfig.Unpack(&scanConfig); err != nil {
							continue
						}
						if !scanConfig.Scan {
							break
						}
						sp, err := time.ParseDuration(scanConfig.ScanPeriod)
						if err != nil {
							continue
						}
						scanPeriod = sp
					}
				}
			}
		}()
	}
}

func Sync() {
	logncore.Sync()
}

func GetLogger(name ...string) core.Logger {
	return logncore.GetLogger(name...)
}
