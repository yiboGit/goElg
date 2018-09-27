package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuffix(t *testing.T) {
	tree := NewSuffixTree("abcc")
	assert.Equal(t, true, tree.HasSubstr("bc"))
}

func TestLsc(t *testing.T) {
	s := findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"})
	assert.Equal(t, "apple", s)
}
