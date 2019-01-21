package rollingfile

import (
	"github.com/shanexu/logn/common"
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
max_size: -1
encoder:
 json:`, true},
		{"case2", `
file_name: /tmp/app.log
encoder:
 json:`, false},
	}

	for _, c := range tests {
		cfg, err := common.NewConfigFrom(c.config)
		assert.Nil(t, err, c.name)
		_, err = NewRollingFile(cfg)
		assert.Equal(t, c.hasErr, err != nil, c.name)
	}
}
