package logn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLogger(t *testing.T) {
	assert.NotNil(t, GetLogger("hello"))
}
