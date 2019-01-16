package rollingfile

import (
	"github.com/shanexu/logp/appender"
	"github.com/spf13/viper"
)

type RollingFile struct {

}

func NewRollingFile(viper viper.Viper) (appender.Appender, error) {
	return &RollingFile{}, nil
}

func init() {
	if err := appender.RegisterAppender("rolling_file", NewRollingFile); err != nil {
		panic(err)
	}
}

