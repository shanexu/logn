package rollingfile

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRollingFile(t *testing.T) {
	cases := []struct{
		config string
		hasErr bool
	}{
		{`
file_name: /tmp/app.log
max_size: -1`, true},
		{`
file_name: /tmp/app.log`, false},
	}

	for _, c := range cases {
		v := viper.New()
		v.SetConfigType("yml")
		err := v.ReadConfig(bytes.NewReader([]byte(c.config)))
		assert.Nil(t, err)
		_, err = NewRollingFile(v)
		assert.Equal(t, c.hasErr, err != nil)
	}
}