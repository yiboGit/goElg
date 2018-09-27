package algorithm

import (
	"container/heap"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *Heap) Pop() interface{} {
	o := *h
	r := o[len(o)-1]
	*h = o[0 : len(o)-1]
	return r
}

type City struct {
	id, costS, stopS int
	index            int
}
type MinPriorityQueue []*City

func (pq MinPriorityQueue) Len() int { return len(pq) }
func (pq MinPriorityQueue) Less(i, j int) bool {
	return pq[i].costS < pq[j].costS
}
func (pq MinPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	// index 不变，其他数据交换
	pq[i].index = i
	pq[j].index = j
}
func (pq *MinPriorityQueue) Push(x interface{}) {
	i := x.(*City)
	i.index = len(*pq)
	*pq = append(*pq, i)
}
func (pq *MinPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *MinPriorityQueue) Update(x *City, costS int) {
	x.costS = costS
	heap.Fix(pq, x.index)
}
