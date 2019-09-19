package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	c := DefaultConfig()
	assert.Empty(t, c.FileName)
}
