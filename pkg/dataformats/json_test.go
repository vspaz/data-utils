package dataformats

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewJsonDecoderOk(t *testing.T) {
	jsonString := `{"foo":10}`
	reader := bufio.NewReader(strings.NewReader(jsonString))
	decoder := NewJsonDecoder(reader)
	var keyToValue map[string]int
	decoder.FromJson(&keyToValue)
	assert.Equal(t, 10, keyToValue["foo"])
}

func TestNewJsonEncoderOk(t *testing.T) {
	out := new(bytes.Buffer)
	encoder := NewJsonEncoder(out)
	encoder.ToJson(map[string]int{"foo": 10})
	assert.Equal(t, "{\n  \"foo\": 10\n}\n", out.String())
}
