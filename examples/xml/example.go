package main

import (
	"encoding/xml"
	"fmt"
	"github.com/vspaz/data-utils/pkg/dataformats"
	"github.com/vspaz/data-utils/pkg/filesystem"
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

func writeXmlFile() {
	dumpFile := filesystem.CreateFile("dump.xml")
	defer filesystem.MustClose(dumpFile)

	encoder := dataformats.NewXmlEncoder(dumpFile)
	users := Users{
		Users: []User{
			{1, "John Doe", "Doe"},
		},
	}
	encoder.ToXml(users)
}

func readXmlFile() {
	fh := filesystem.OpenFile("dump.xml")
	defer filesystem.MustClose(fh)

	decoder := dataformats.NewXmlDecoder(fh)
	users := &Users{}
	decoder.FromXml(users)
	fmt.Println(users.Users[0].Name)
}

func main() {
	writeXmlFile()
	readXmlFile()
}
