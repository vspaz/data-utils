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

func NewRowReader(in io.Reader, delimiter string) Iterator {
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

type Writer struct {
	writer *csv.Writer
}

func (w *Writer) write(values ...string) {
	err := w.writer.Write(values)
	if err != nil {
		log.Printf("error to write to file %s '%v'", err.Error(), values)
	}
}

func NewRowWriter(out io.Writer) *Writer {
	writer := csv.NewWriter(out)
	return &Writer{writer: writer}
}
