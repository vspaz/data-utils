# data-utils

A small simple-to-use Go module for working with various data formats.

## Delimited text files (CSV, TSV etc.)

### Writing to a CSV file

```go
package main

import (
	"fmt"
	"github.com/vspaz/data-utils/pkg/dataformats"
	"github.com/vspaz/data-utils/pkg/filesystem"
)

func writeTsvFile() {
	dumpFile := filesystem.CreateFile("dump.tsv")
	defer filesystem.MustClose(dumpFile)

	csvWriter := dataformats.NewRowWriter(dumpFile, ",") // any of {"\t", ",", " ", ";", "|"}
	csvWriter.Write("value1", "value2", "value3")
	csvWriter.Write("value4", "value5", "value6")
	csvWriter.Flush()
}

```

### Reading from a CSV file

```go
package main

import (
	"fmt"
	"github.com/vspaz/data-utils/pkg/dataformats"
	"github.com/vspaz/data-utils/pkg/filesystem"
)

func readTsvFile() {
	fh := filesystem.OpenFile("dump.tsv")
	defer filesystem.MustClose(fh)

	csvReader := dataformats.NewRowReader(fh, ",") // any of {"\t", ",", " ", ";", "|"}
	for csvReader.HasNext() {
		row := csvReader.Next()
		fmt.Println(row[0], row[1], row[2])
	}
}

```

### Writing to a yaml file

```go
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

```

### Reading from a yaml file

```go
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

func readYamlFile() {
	fh := filesystem.OpenFile("dump.yaml")
	defer filesystem.MustClose(fh)

	decoder := dataformats.NewYamlDecoder(fh)
	data := &MyStruct{}
	decoder.FromYaml(data)
	fmt.Println(data.Content)
}

```

### Writing to a json file

```go
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
```

### Reading from a json file

```go
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

func readJsonFile() {
	fh := filesystem.OpenFile("dump.json")
	defer filesystem.MustClose(fh)

	decoder := dataformats.NewJsonDecoder(fh)
	data := &MyStruct{}
	decoder.FromJson(data)
	fmt.Println(data.Content)
}

```
