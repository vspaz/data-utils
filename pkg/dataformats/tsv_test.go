package dataformats

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewRowIteratorTsvOk(t *testing.T) {
	tsvRecords := `
foo1	bar1	baz1
foo2	bar2	baz2
`
	records := bufio.NewReader(strings.NewReader(tsvRecords))
	iterator := NewRowIterator(records, "tsv")

	assert.True(t, iterator.HasNext())
	record1 := iterator.Next()
	assert.Equal(t, []string{"foo1", "bar1", "baz1"}, record1)

	assert.True(t, iterator.HasNext())
	record2 := iterator.Next()
	assert.Equal(t, []string{"foo2", "bar2", "baz2"}, record2)
}

func TestNewRowIteratorCsvOk(t *testing.T) {
	csvRecords := `
foo1,bar1,baz1
foo2,bar2,baz2
`
	records := bufio.NewReader(strings.NewReader(csvRecords))
	iterator := NewRowIterator(records, "csv")

	assert.True(t, iterator.HasNext())
	record1 := iterator.Next()
	assert.Equal(t, []string{"foo1", "bar1", "baz1"}, record1)

	assert.True(t, iterator.HasNext())
	record2 := iterator.Next()
	assert.Equal(t, []string{"foo2", "bar2", "baz2"}, record2)
}
