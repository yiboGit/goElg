package algorithm

import (
	"log"
	"strings"
)

type Node struct {
	key       string
	childrens map[string]*Node
	IsLeaf    bool
	value     int
}
type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	root := &Node{
		key:       "/",
		childrens: make(map[string]*Node),
		IsLeaf:    false,
	}
	return &Trie{root}
}

func (t *Trie) AddPath(path string, v int) {
	arr := strings.Split(path, "/")
	l := len(arr)
	if l == 0 {
		return
	}
	t.addNode(t.root, arr[1:], v)
}

func (t *Trie) Print(p *Node) {
	for k, v := range p.childrens {
		log.Printf("k: %s, v %v ", k, *v)
		t.Print(p.childrens[k])
	}
}
func (t *Trie) addNode(node *Node, pathSlice []string, v int) {
	if len(pathSlice) == 0 {
		node.IsLeaf = true
		node.value = v
		return
	}
	leftMost := pathSlice[0]
	subNode, ok := node.childrens[leftMost]
	if !ok {
		subNode = &Node{
			key:       node.key + leftMost + "/",
			childrens: make(map[string]*Node),
			IsLeaf:    false,
		}
		node.childrens[leftMost] = subNode
	}
	subNode.IsLeaf = false
	t.addNode(subNode, pathSlice[1:], v)
}

func (t *Trie) Match(path string) int {
	arr := strings.Split(path, "/")
	return t.Find(t.root, arr[1:])
}
func (t *Trie) Find(node *Node, pathSlice []string) int {
	if len(pathSlice) == 0 {
		return node.value
	}
	subNote, ok := node.childrens[pathSlice[0]]
	if !ok {
		return -1
	}
	return t.Find(subNote, pathSlice[1:])
}
