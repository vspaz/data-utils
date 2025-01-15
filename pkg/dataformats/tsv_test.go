package dataformats

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewRowIteratorOk(t *testing.T) {
	t.Parallel()
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

func BenchmarkNewRowIterator(b *testing.B) {
	csvRecords := `
1,foo,bar,baz
2,foo,bar,baz
3,foo,bar,baz
4,foo,bar,baz
5,foo,bar,baz
6,foo,bar,baz
7,foo,bar,baz
8,foo,bar,baz
9,foo,bar,baz
10,foo,bar,baz
`
	records := bufio.NewReader(strings.NewReader(csvRecords))
	for i := 0; i < b.N; i++ {
		iterator := NewRowReader(records, ",")
		for iterator.HasNext() {
			_ = iterator.Next()
		}
	}
}

func TestNewRowIteratorTsvOk(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
	out := new(bytes.Buffer)
	writer := NewRowWriter(out, "\t")
	writer.Write("foo1", "bar1", "baz1")
	writer.Write("foo2", "bar2", "baz2")
	writer.Flush()
	record := out.String()
	assert.Equal(t, "foo1\tbar1\tbaz1\nfoo2\tbar2\tbaz2\n", record)
}

func TestNewRowWriterCsvOk(t *testing.T) {
	t.Parallel()
	out := new(bytes.Buffer)
	writer := NewRowWriter(out, ",")
	writer.Write("foo", "bar", "baz")
	writer.Flush()
	record := out.String()
	assert.Equal(t, "foo,bar,baz\n", record)
}

func TestNewRowWriterTsvOk(t *testing.T) {
	t.Parallel()
	out := new(bytes.Buffer)
	writer := NewRowWriter(out, "\t")
	writer.Write("foo", "bar", "baz")
	writer.Flush()
	record := out.String()
	assert.Equal(t, "foo\tbar\tbaz\n", record)
}
