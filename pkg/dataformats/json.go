package dataformats

import (
	"encoding/json"
	"io"
	"log"
)

type JsonDecoder struct {
	decoder *json.Decoder
}

func NewJsonDecoder(in io.Reader) *JsonDecoder {
	return &JsonDecoder{decoder: json.NewDecoder(in)}
}

func (j *JsonDecoder) FromJson(deserializable any) {
	if err := j.decoder.Decode(deserializable); err != nil && err != io.EOF {
		log.Fatalf("error decoding json %s", err)
	}
}

type JsonEncoder struct {
	encoder *json.Encoder
}

func NewJsonEncoder(out io.Writer) *JsonEncoder {
	return &JsonEncoder{encoder: json.NewEncoder(out)}
}

func (j *JsonEncoder) ToJson(serializable any) {
	err := j.encoder.Encode(serializable)
	if err != nil {
		log.Fatalf("error encoding json %s", err)
	}
}
