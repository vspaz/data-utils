package dataformats

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
)

type YamlDecoder struct {
	decoder *yaml.Decoder
}

func NewYamlDecoder(in io.Reader) *YamlDecoder {
	return &YamlDecoder{decoder: yaml.NewDecoder(in)}
}

func (y *YamlDecoder) FromYaml(deserializable any) {
	if err := y.decoder.Decode(deserializable); err != nil && err != io.EOF {
		log.Fatalf("error decoding yaml %s", err)
	}
}

type YamlEncoder struct {
	encoder *yaml.Encoder
}

func NewYamlEncoder(out io.Writer) *YamlEncoder {
	return &YamlEncoder{encoder: yaml.NewEncoder(out)}
}

func (y *YamlEncoder) ToYaml(serializable any) {
	err := y.encoder.Encode(serializable)
	if err != nil {
		log.Fatalf("error encoding yaml %s", err)
	}
}
