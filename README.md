# data-utils
A small Go module for working with various data formats.

## CSV

### Writing to TSV file

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

	csvWriter := dataformats.NewRowWriter(dumpFile, ",")  // any of {"\t", ",", " ", ";", "|"}
	csvWriter.Write("value1", "value2", "value3")
	csvWriter.Write("value4", "value5", "value6")
	csvWriter.Flush()
}

func readTsvFile() {
	fh := filesystem.OpenFile("dump.tsv")
	defer filesystem.MustClose(fh)

	csvReader := dataformats.NewRowReader(fh, ",")  // any of {"\t", ",", " ", ";", "|"}
	for csvReader.HasNext() {
		row := csvReader.Next()
		fmt.Println(row[0], row[1], row[2])
	}
}

func main() {
	writeTsvFile()
	readTsvFile()
}
```

