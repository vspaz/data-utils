package dataformats

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

type User struct {
	ID    int    `xml:"id,attr"`
	Login string `xml:"login"`
	Name  string `xml:"name"`
}

var xmlText = `<?xml version="1.0" encoding="UTF-8"?>
	<users>
		<user id="138">
			<login>johndoe</login>
			<name>Doe</name>
		</user>
	</users>`

func TestXmlDecoderOk(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(xmlText))
	decoder := NewXmlDecoder(reader)
	users := Users{}
	decoder.FromXml(&users)
	assert.Equal(t, 138, users.Users[0].ID)
}

func TestXmlEncoderOk(t *testing.T) {
	out := new(bytes.Buffer)
	encoder := NewXmlEncoder(out)
	encoder.toXml(User{ID: 138, Login: "johndoe", Name: "Doe"})
	assert.Equal(t, "<User id=\"138\">\n <login>johndoe</login>\n <name>Doe</name>\n</User>", out.String())
}
