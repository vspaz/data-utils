package dataformats

import (
	"encoding/csv"
)

type Iterator interface {
	Next() []string
	HasNext() bool
}

type RowIterator struct {
	csvReader  *csv.Reader
	fileFormat string // csv | tsv
	currentRow []string
}

func (r *RowIterator) Next() []string {
	return r.currentRow
}

func (r *RowIterator) HasNext() bool {
	record, err := r.csvReader.Read()
	if err != nil {
		return false
	}
	r.currentRow = record
	return true
}

func NewRowIterator(reader *csv.Reader, fileFormat string) Iterator {
	if fileFormat == "tsv" {
		reader.Comma = '\t'
	}
	return &RowIterator{
		csvReader:  reader,
		fileFormat: fileFormat,
	}
}
