package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shanexu/logn/common"
	"github.com/shanexu/logn/core/zap"
	_ "github.com/shanexu/logn/includes"
)

func TestNew(t *testing.T) {
	rawConfig, err := common.NewConfigFrom(`
appenders:
  console:
    - name: CONSOLE
      target: stdout
      encoder:
        json:
  file:
    - name: FILE
      file_name: /tmp/app.log
      encoder:
        json:
    - name: METRICS
      file_name: /tmp/metrics.log
      encoder:
        json:
loggers:
  root:
    level: info
    appender_refs:
      - CONSOLE
  logger:
    - name: helloworld
      appender_refs:
        - CONSOLE
        - FILE
      level: debug
`)
	if err != nil {
		t.Fatal(err)
	}
	c, err := zap.New(rawConfig)
	assert.NotNil(t, c)
	assert.Nil(t, err)

	hello := c.GetLogger("hello")
	hello.Info("hello")
	world := c.GetLogger("world")
	world.Debug("world")

	helloworld := c.GetLogger("helloworld")
	helloworld.Info("hello world")
	c.Sync()
}
