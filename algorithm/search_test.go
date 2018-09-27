package algorithm

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	r := SearchInSortedRotateArray([]int{4, 5, 1, 2, 3}, 3)
	assert.Equal(t, 4, r)
}

func TestQsort(t *testing.T) {
	o := []int{1, 3, 0, 2, 4, 5}
	qsort(o)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, o)
}
func TestTopk(t *testing.T) {
	o := []int{1, 3, 0, 2, 4, 5}
	d := TopK(o, 5)
	assert.Equal(t, 4, d)
}

func TestShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Shuffle(arr)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
}

func TestSubArray(t *testing.T) {
	arr := []int{1, -1, 9, -1, 0, 2, -2}
	r := MaxSumOfSubArray(arr)
	assert.Equal(t, 10, r)
}

func TestLongestPandir(t *testing.T) {
	s := "xabbax"
	r := LongestPandir(s)
	assert.Equal(t, "xabbax", r)
}

func TestBSstart(t *testing.T) {
	r := BSstart([]int{1, 1, 3}, 1)
	// a := []int{1, 2, 3}
	assert.Equal(t, 4, r)
}

func TestCutRod(t *testing.T) {
	m := CutRod([]int{1, 5, 8, 9, 10, 17, 17, 20}, 8)
	assert.Equal(t, 22, m)
}

func TestFindMin(t *testing.T) {
	arr := []int{1}
	assert.Equal(t, 1, findMin(arr))
	arr = []int{1, 2}
	assert.Equal(t, 1, findMin(arr))
	arr = []int{1, 2, 3}
	assert.Equal(t, 1, findMin(arr))
	arr = []int{2, 1, 3, 0, 5}
	assert.Equal(t, 0, findMin(arr))
}

func TestkthSmallest(t *testing.T) {

}

func TestPermute(t *testing.T) {
	Permute("abcd")
	assert.Equal(t, 0, 1)
}

func TestBuildTree(t *testing.T) {
	arr := []int{1, 2, 4, 5}
	node := BuildSegmentTree(arr, 0, len(arr)-1)
	// printTree(node)
	sum := QueryRange(arr, node, 2, 5)
	assert.Equal(t, 9, sum)
}

func TestHeap(t *testing.T) {
	h := &Heap{3, 2, 1}
	heap.Init(h)
}

func TestPriority(t *testing.T) {
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	pq := make(MinPriorityQueue, len(items))
	i := 0
	for _, priority := range items {
		pq[i] = &City{id: i + 1, costS: priority}
		i++
	}
	heap.Init(&pq)

	item := &City{
		id:    0,
		costS: 6,
	}
	heap.Push(&pq, item)
	pq.Update(item, 0)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*City)
		fmt.Printf("%d:%d \n", item.id, item.costS)
	}
}

func TestRotateMatrix(t *testing.T) {
	M := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}
	RotatePrintMatrix(M)
}
