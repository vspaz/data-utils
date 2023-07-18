package dataformats

import (
	"bytes"
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

func TestNewYamlEncoderOk(t *testing.T) {
	type TestObject struct {
		Foo int      `yaml:"foo"`
		Bar []string `yaml:"bar"`
	}
	tempHolder := &TestObject{
		Foo: 10,
		Bar: []string{"foo", "bar", "baz"},
	}
	out := new(bytes.Buffer)
	encoder := NewYamlEncoder(out)
	encoder.ToYaml(tempHolder)
	assert.Equal(t, "foo: 10\nbar:\n  - foo\n  - bar\n  - baz\n", out.String())
}
