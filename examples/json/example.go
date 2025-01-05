package main

import (
	"fmt"
	"github.com/vspaz/data-utils/pkg/dataformats"
	"github.com/vspaz/data-utils/pkg/filesystem"
)

type MyStruct struct {
	Content string            `json:"content,omitempty"`
	Config  map[string]string `json:"config,omitempty"`
	Params  []string          `json:"params,omitempty"`
}

func writeJsonFile() {
	dumpFile := filesystem.CreateFile("dump.json")
	defer filesystem.MustClose(dumpFile)

	encoder := dataformats.NewJsonEncoder(dumpFile)
	data := MyStruct{
		Content: "some content goes here",
		Config:  map[string]string{"foo": "bar"},
		Params:  []string{"value1", "value2", "value3"},
	}
	encoder.ToJson(&data)
}

func readJsonFile() {
	fh := filesystem.OpenFile("dump.json")
	defer filesystem.MustClose(fh)

	decoder := dataformats.NewJsonDecoder(fh)
	data := &MyStruct{}
	decoder.FromJson(data)
	fmt.Println(data.Content)
}

func main() {
	writeJsonFile()
	readJsonFile()
}
