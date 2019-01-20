package core

import (
	"errors"
	"github.com/shanexu/logp/appender"
	"github.com/shanexu/logp/common"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Core struct {
	nameToLogger   sync.Map
	nameToAppender map[string]appender.Appender
	rootAppender   *appender.Appender
	rootLevel      zapcore.Level
}

func New(rawConfig *common.Config) (*Core, error) {
	rawConfig.Unpack()

	return nil, errors.New("not implemented")
}
