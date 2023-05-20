package dataformats

import (
	"encoding/csv"
	"io"
	"log"
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
	switch delimiter {
	case "\t":
		reader.Comma = '\t'
	case ",":
		reader.Comma = ','
	case " ":
		reader.Comma = ' '
	case ";":
		reader.Comma = ';'
	case "|":
		reader.Comma = '|'
	default:
		log.Panicf("invalid delimiter '%v'; '\\t', ',', ' ', ';', '|' are only allowed", delimiter)
	}
	return &RowIterator{
		reader:    reader,
		delimiter: delimiter,
	}
}
