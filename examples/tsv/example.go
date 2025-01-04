package main

import (
	"fmt"
	"github.com/vspaz/data-utils/pkg/dataformats"
	"github.com/vspaz/data-utils/pkg/filesystem"
)

func writeTsvFile() {
	dumpFile := filesystem.CreateFile("dump.tsv")
	defer filesystem.MustClose(dumpFile)

	csvWriter := dataformats.NewRowWriter(dumpFile, ",")
	csvWriter.Write("value1", "value2")
	csvWriter.Write("value3", "value4")
	csvWriter.Flush()
}

func readTsvFile() {
	fh := filesystem.OpenFile("dump.tsv")
	defer filesystem.MustClose(fh)

	csvReader := dataformats.NewRowReader(fh, ",")
	for csvReader.HasNext() {
		row := csvReader.Next()
		fmt.Println(row[0], row[1])
	}
}

func main() {
	writeTsvFile()
	readTsvFile()
}
