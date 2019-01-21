package file

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	c := DefaultConfig()
	assert.Empty(t, c.FileName)
}