package dataformats

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var tsvRecords = `
foo1\tbar1\tbaz1
foo2\tbar2\tbaz2
`

func TestNewRowIterator(t *testing.T) {
	records := bufio.NewReader(strings.NewReader(tsvRecords))
	iterator := NewRowIterator(records, "tsv")
	assert.True(t, iterator.HasNext())
}
