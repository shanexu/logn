package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	assert.Equal(t, "info", cfg.Loggers.Root.Level)
	assert.Equal(t, 0, len(cfg.Loggers.Logger))
	assert.Equal(t, 1, len(cfg.Appenders))
	assert.Equal(t, 1, len(cfg.Appenders["console"]))
	consoleCfg := cfg.Appenders["console"][0]
	name, _ := consoleCfg.String("name", -1)
	assert.Equal(t, "CONSOLE", name)
	target, _ := consoleCfg.String("target", -1)
	assert.Equal(t, "stdout", target)
	encoder, _ := consoleCfg.Child("encoder", -1)
	assert.NotNil(t, encoder)
	assert.True(t, encoder.HasField("json"))
}
