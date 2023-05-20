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
		log.Fatalf("error decoding yaml %s", err)
	}
}

func FromJson(deserializable []byte, deserialized any) error {
	return json.Unmarshal(deserializable, deserialized)
}

func ToJson(serializable any) []byte {
	encodedMessage, _ := json.Marshal(serializable)
	return encodedMessage
}
