package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	tr := NewTrie()
	tr.AddPath("/a", 1)
	tr.AddPath("/a/b", 2)
	tr.AddPath("/b", 3)
	tr.AddPath("/a/b/c", 4)
	assert.Equal(t, 1, tr.Match("/a"))
	assert.Equal(t, 2, tr.Match("/a/b"))
	assert.Equal(t, 3, tr.Match("/b"))
	assert.Equal(t, 4, tr.Match("/a/b/c"))
}
