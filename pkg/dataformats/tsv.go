package dataformats

import (
	"encoding/csv"
	"io"
)

type Iterator interface {
	Next() []string
	HasNext() bool
}

type RowIterator struct {
	reader     *csv.Reader
	delimiter  string // csv | tsv
	currentRow []string
}

func (r *RowIterator) Next() []string {
	return r.currentRow
}

func (r *RowIterator) HasNext() bool {
	row, err := r.reader.Read()
	if err != nil {
		return false
	}
	r.currentRow = row
	return true
}

func NewRowIterator(in io.Reader, delimiter string) Iterator {
	reader := csv.NewReader(in)
	if delimiter == "tsv" {
		reader.Comma = '\t'
	}
	return &RowIterator{
		reader:    reader,
		delimiter: delimiter,
	}
}
