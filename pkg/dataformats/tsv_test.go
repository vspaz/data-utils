package dataformats

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var tsvRecords = `
foo1	bar1	baz1
foo2	bar2	baz2
`

func TestNewRowIterator(t *testing.T) {
	records := bufio.NewReader(strings.NewReader(tsvRecords))
	iterator := NewRowIterator(records, "tsv")

	assert.True(t, iterator.HasNext())
	assert.Equal(t, "foo1", iterator.Next()[0])
	assert.Equal(t, "bar1", iterator.Next()[1])
	assert.Equal(t, "baz1", iterator.Next()[2])

	assert.Equal(t, iterator.Next(), []string{"foo2", "bar2", "baz2"})
}
