package dataformats

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"io"
	"log"
)

type JsonDecoder struct {
	decoder *jsontext.Decoder
}

func NewJsonDecoder(in io.Reader) *JsonDecoder {
	return &JsonDecoder{decoder: jsontext.NewDecoder(in)}
}

func (j *JsonDecoder) FromJson(deserializable any) {
	if err := json.UnmarshalDecode(j.decoder, &deserializable); err != nil && err != io.EOF {
		log.Fatalf("error decoding json %s", err)
	}
}

type JsonEncoder struct {
	encoder *jsontext.Encoder
}

func NewJsonEncoder(out io.Writer, indent string, escapeHtml bool) *JsonEncoder {
	encoder := jsontext.NewEncoder(
		out,
		jsontext.WithIndent(indent),
		jsontext.EscapeForHTML(escapeHtml),
	)
	return &JsonEncoder{
		encoder: encoder,
	}
}

func (j *JsonEncoder) ToJson(serializable any) {
	err := json.MarshalEncode(
		j.encoder,
		serializable,
	)
	if err != nil {
		log.Fatalf("error encoding json %s", err)
	}
}
