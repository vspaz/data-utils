package dataformats

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewRowIteratorOk(t *testing.T) {
	csvRecords := `
foo,bar,baz
foo,bar,baz
foo,bar,baz
`
	records := bufio.NewReader(strings.NewReader(csvRecords))
	iterator := NewRowReader(records, ",")
	recordCount := 0
	for iterator.HasNext() {
		record := iterator.Next()
		assert.Equal(t, []string{"foo", "bar", "baz"}, record)
		recordCount++
	}
	assert.Equal(t, 3, recordCount)
}

func TestNewRowIteratorTsvOk(t *testing.T) {
	tsvRecords := `
foo1	bar1	baz1
foo2	bar2	baz2
`
	records := bufio.NewReader(strings.NewReader(tsvRecords))
	iterator := NewRowReader(records, "\t")

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
	iterator := NewRowReader(records, ",")

	assert.True(t, iterator.HasNext())
	record1 := iterator.Next()
	assert.Equal(t, []string{"foo1", "bar1", "baz1"}, record1)

	assert.True(t, iterator.HasNext())
	record2 := iterator.Next()
	assert.Equal(t, []string{"foo2", "bar2", "baz2"}, record2)
}

func TestHasNextOk(t *testing.T) {
	csvRecords := `foo,bar,baz`
	records := bufio.NewReader(strings.NewReader(csvRecords))
	iterator := NewRowReader(records, ",")
	recordCount := 0
	for iterator.HasNext() {
		record := iterator.Next()
		assert.Equal(t, []string{"foo", "bar", "baz"}, record)
		recordCount++
	}
	assert.False(t, iterator.HasNext())
	assert.Equal(t, 1, recordCount)
}

func TestNewRowWriterOk(t *testing.T) {
	out := new(bytes.Buffer)
	writer := NewRowWriter(out, "\t")
	writer.write("foo1", "bar1", "baz1")
	writer.write("foo2", "bar2", "baz2")
	writer.Flush()
	record := out.String()
	assert.Equal(t, "foo1\tbar1\tbaz1\nfoo2\tbar2\tbaz2\n", record)
}

func TestNewRowWriterCsvOk(t *testing.T) {
	out := new(bytes.Buffer)
	writer := NewRowWriter(out, ",")
	writer.write("foo", "bar", "baz")
	writer.Flush()
	record := out.String()
	assert.Equal(t, "foo,bar,baz\n", record)
}

func TestNewRowWriterTsvOk(t *testing.T) {
	out := new(bytes.Buffer)
	writer := NewRowWriter(out, "\t")
	writer.write("foo", "bar", "baz")
	writer.Flush()
	record := out.String()
	assert.Equal(t, "foo\tbar\tbaz\n", record)
}
