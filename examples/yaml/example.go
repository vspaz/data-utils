package main

import (
	"fmt"
	"github.com/vspaz/data-utils/pkg/dataformats"
	"github.com/vspaz/data-utils/pkg/filesystem"
)

type MyStruct struct {
	Content string            `yaml:"content,omitempty"`
	Config  map[string]string `yaml:"config,omitempty"`
	Params  []string          `yaml:"params,omitempty"`
}

func writeYamlFile() {
	dumpFile := filesystem.CreateFile("dump.yaml")
	defer filesystem.MustClose(dumpFile)

	encoder := dataformats.NewYamlEncoder(dumpFile)
	data := MyStruct{
		Content: "some content goes here",
		Config:  map[string]string{"foo": "bar"},
		Params:  []string{"value1", "value2", "value3"},
	}
	encoder.ToYaml(&data)
}

func readYamlFile() {
	fh := filesystem.OpenFile("dump.yaml")
	defer filesystem.MustClose(fh)

	decoder := dataformats.NewYamlDecoder(fh)
	someStruct := &MyStruct{}
	decoder.FromYaml(someStruct)
	fmt.Println(someStruct.Content)
}

func main() {
	writeYamlFile()
	readYamlFile()
}
