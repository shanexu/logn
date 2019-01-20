package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNamedConfig(t *testing.T)  {
	cfg, err := NewConfigFrom(`
name: hello
a: 1
b: 2
`)
	if err != nil {
		t.Fatal(err)
	}
	nc := NamedConfig{}
	err = cfg.Unpack(&nc)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "hello", nc.Name)
	c, _ := NewConfigFrom(nc.Config)
	fmt.Println(c.Int("a", -1))
	fmt.Println(c.Int("b", -1))
}