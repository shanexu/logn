package rollingfile

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRollingFile(t *testing.T) {
	tests := []struct {
		name   string
		config string
		hasErr bool
	}{
		{"case1", `
file_name: /tmp/app.log
max_size: -1`, true},
		{"case2", `
file_name: /tmp/app.log`, false},
	}

	for _, c := range tests {
		v := viper.New()
		v.SetConfigType("yml")
		err := v.ReadConfig(bytes.NewReader([]byte(c.config)))
		assert.Nil(t, err)
		_, err = NewRollingFile(v)
		assert.Equal(t, c.hasErr, err != nil)
	}
}
