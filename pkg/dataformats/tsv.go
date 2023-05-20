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

type Reader struct {
	reader     *csv.Reader
	currentRow []string
}

func (r *Reader) Next() []string {
	return r.currentRow
}

func (r *Reader) HasNext() bool {
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
	return &Reader{reader: reader}
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

func (w *Writer) Flush() {
	w.writer.Flush()
}

func NewRowWriter(out io.Writer, delimiter string) *Writer {
	writer := csv.NewWriter(out)
	switch delimiter {
	case "\t":
		writer.Comma = '\t'
	case ",":
		writer.Comma = ','
	case " ":
		writer.Comma = ' '
	case ";":
		writer.Comma = ';'
	case "|":
		writer.Comma = '|'
	default:
		log.Panicf("invalid delimiter '%v'; '\\t', ',', ' ', ';', '|' are only allowed", delimiter)
	}
	return &Writer{writer: writer}
}
