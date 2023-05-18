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

func (r RowIterator) Next() {
	//TODO implement me
	panic("implement me")
}

func (r RowIterator) HasNext() {
	//TODO implement me
	panic("implement me")
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
