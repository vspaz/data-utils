# data-utils

A small simple-to-use Go module for working with various data formats.

## Delimited text files (CSV, TSV etc.)

### Writing to an CSV file

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

### Reading a CSV file

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

### Reading a yaml file

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

### Reading a json file

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

### Writing to an XML file

```go
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
```

### Reading an XML file

```go
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

func readXmlFile() {
	fh := filesystem.OpenFile("dump.xml")
	defer filesystem.MustClose(fh)

	decoder := dataformats.NewXmlDecoder(fh)
	users := &Users{}
	decoder.FromXml(users)
	fmt.Println(users.Users[0].Name)
}
```