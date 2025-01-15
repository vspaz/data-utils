package dataformats

import (
	"encoding/xml"
	"io"
	"log"
)

type XmlDecoder struct {
	decoder *xml.Decoder
}

func NewXmlDecoder(in io.Reader) *XmlDecoder {
	return &XmlDecoder{decoder: xml.NewDecoder(in)}
}

func (x *XmlDecoder) FromXml(deserializable any) {
	if err := x.decoder.Decode(deserializable); err != nil {
		log.Fatalf("error decoding XML: %s", err)
	}
}

type XmlEncoder struct {
	encoder *xml.Encoder
}

func NewXmlEncoder(out io.Writer) *XmlEncoder {
	encoder := xml.NewEncoder(out)
	encoder.Indent("", " ")
	return &XmlEncoder{encoder: encoder}
}

func (x *XmlEncoder) toXml(serializable any) {
	err := x.encoder.Encode(serializable)
	if err != nil {
		log.Fatalf("error encoding XML: %s", err)
	}
}
