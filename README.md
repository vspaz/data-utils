# data-utils
A small Go module for working with various data formats.

## Delimited text files (CSV, TSV etc.)

### Writing to CSV file

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

	csvReader := dataformats.NewRowReader(fh, ",")  // any of {"\t", ",", " ", ";", "|"}
	for csvReader.HasNext() {
		row := csvReader.Next()
		fmt.Println(row[0], row[1], row[2])
	}
}

```

