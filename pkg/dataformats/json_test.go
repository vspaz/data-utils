package dataformats

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewJsonDecoderOk(t *testing.T) {
	jsonString := `{"foo":"bar"}`
	reader := bufio.NewReader(strings.NewReader(jsonString))
	decoder := NewJsonDecoder(reader)
	var keyToValue map[string]string
	decoder.FromJson(&keyToValue)
	assert.Equal(t, "bar", keyToValue["foo"])
}
