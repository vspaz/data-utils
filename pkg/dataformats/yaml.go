package dataformats

import (
	"gopkg.in/yaml.v3"
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

type CustomDate string

// this is a custom yaml date murshaller to avoid wrapping strings that contain characters
// :, {, }, [, ], ,, &, *, #, ?, |, -, <, >, =, !, %, @, ` with double qoutes:
func (c CustomDate) MarshalYAML() (any, error) {
	return yaml.Node{
		Kind:  yaml.ScalarNode,
		Style: yaml.FlowStyle,
		Value: string(c),
	}, nil
}

func NewYamlEncoder(out io.Writer) *YamlEncoder {
	encoder := yaml.NewEncoder(out)
	encoder.SetIndent(2)
	return &YamlEncoder{encoder: encoder}
}

func (y *YamlEncoder) ToYaml(serializable any) {
	err := y.encoder.Encode(serializable)
	if err != nil {
		log.Fatalf("error encoding yaml %s", err)
	}
}
