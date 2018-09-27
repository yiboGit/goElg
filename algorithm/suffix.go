package algorithm

import (
	"log"
)

type SuffixNode struct {
	// parent-c-children
	Children   map[string]*SuffixNode
	SuffixLink *SuffixNode
}

func (sn *SuffixNode) AddChildren(c string, child *SuffixNode) {
	sn.Children[c] = child

}
func NewSuffixNode(link *SuffixNode) *SuffixNode {
	node := &SuffixNode{Children: make(map[string]*SuffixNode)}
	if link != nil {
		node.SuffixLink = link
	} else {
		node.SuffixLink = node
	}
	return node
}

type SuffixTree struct {
	Root *SuffixNode
}

func NewSuffixTree(str string) *SuffixTree {
	if len(str) == 0 {
		return nil
	}
	root := NewSuffixNode(nil)
	tree := &SuffixTree{root}
	tree.Build(str)
	return tree
}

func (s *SuffixTree) Build(str string) {
	longest := NewSuffixNode(s.Root)
	s.Root.AddChildren(string(str[0]), longest)
	for _, c := range str[1:] {
		k := string(c)
		curr := longest
		var prev *SuffixNode
		for {
			_, ok := curr.Children[k]
			log.Printf("not ok: %s", k)
			if ok {
				break
			}
			r1 := NewSuffixNode(nil)
			curr.AddChildren(k, r1)

			if prev != nil {
				prev.SuffixLink = r1
			}
			prev = r1
			curr = curr.SuffixLink
		}
		if curr == s.Root {
			prev.SuffixLink = s.Root
		} else {
			prev.SuffixLink = curr.Children[k]
		}
		longest = longest.Children[k]
	}
	// s.Print(s.Root, 1)
}

func (s *SuffixTree) Print(t *SuffixNode, level int) {
	for k, v := range t.Children {
		log.Printf("k: %s, v %v level %d", k, *v, level)
		s.Print(t.Children[k], level+1)
	}
}

// HasSubstr 是否存在子串
func (s *SuffixTree) HasSubstr(ss string) bool {
	s.Print(s.Root, 1)
	r := s.Root
	for _, c := range ss {
		child, ok := r.Children[string(c)]
		if !ok {
			return false
		}
		r = child
	}
	return true
}
