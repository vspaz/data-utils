package dataformats

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
)

type YamlReader struct {
	reader yaml.Decoder
}

func NewYamlDecoder(in io.Reader) *yaml.Decoder {
	return yaml.NewDecoder(in)
}

func (y *YamlReader) FromYaml(body any) {
	if err := y.reader.Decode(body); err != nil && err != io.EOF {
		log.Fatalf("error decoding yaml %s", err)
	}
}
