package dataformats

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewYamlDecoderOk(t *testing.T) {
	testObject := `---
foo: 3
bar:
   - one
   - two
   - three
   - four
`
	type TestObject struct {
		Foo int      `yaml:"foo"`
		Bar []string `yaml:"bar"`
	}
	decoder := NewYamlDecoder(strings.NewReader(testObject))
	tempHolder := &TestObject{}
	decoder.FromYaml(tempHolder)
	assert.Equal(t, 3, tempHolder.Foo)
}
