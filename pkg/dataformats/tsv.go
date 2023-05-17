package dataformats

import "encoding/csv"

type Iterator interface {
	Next()
	HasNext()
}

type RowIterator struct {
	csvReader  *csv.Reader
	fileFormat string // csv | tsv
	currentRow []string
}
