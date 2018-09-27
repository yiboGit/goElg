package algorithm

import (
	"container/list"
	"fmt"
)

type Stack struct {
	List *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}
func (s *Stack) Push(a int) {
	s.List.PushFront(a)
}
func (s *Stack) Pop(a int) {
	f := s.List.Front()
	s.List.Remove(f)
}
func (s *Stack) Print() {
	f := s.List.Front()
	for f != nil {
		fmt.Print(f.Value.(int))
		f = f.Next()
	}
}

// Combination Sum no duplicate

func combinationSum(arr []int, target int) [][]int {
	cur := make([]int, 0)
	res := make([][]int, 0)
	DfsSum(arr, 0, 0, target, cur, &res)
	fmt.Print(res)
	return res
}

func DfsSum(arr []int, i, sum int, target int, cur []int, res *[][]int) {
	if sum > target {
		return
	}
	if sum == target {
		dst := make([]int, len(cur))
		copy(dst, cur)
		*res = append(*res, dst)
		return
	}
	for j := i; j < len(arr); j++ {
		cur = append(cur, arr[j])
		DfsSum(arr, j, sum+arr[j], target, cur, res)
		cur = cur[:len(cur)-1]
	}
}

// BSstart find leftmost index of sorted array
func BSstart(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return helper(head, nil)
}

func helper(head, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	c := &TreeNode{Val: slow.Val}
	c.Left = helper(head, slow)
	c.Right = helper(slow.Next, tail)
	return c
}

func Lcs(s1, s2 string) int {
	// naive resursive
	return LcsDp(s1, s2, len(s1), len(s2))
}

func LcsHelper(s1, s2 string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if s1[m-1] == s2[n-1] {
		return 1 + LcsHelper(s1, s2, m-1, n-1)
	}
	return max(LcsHelper(s1, s2, m, n-1), LcsHelper(s1, s2, m-1, n))
}

func LcsDp(s1, s2 string, m, n int) int {
	tb := make([][]int, m+1)
	for i, _ := range tb {
		tb[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				tb[i][j] = 0
			} else if s1[i-1] == s2[j-1] {
				tb[i][j] = tb[i-1][j-1] + 1
			} else {
				tb[i][j] = max(tb[i-1][j], tb[i][j-1])
			}
		}
	}
	return tb[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLongestWord(s string, d []string) string {
	lg := ""
	for _, st := range d {
		isContain := findLongestHelper(s, st)
		if !isContain {
			continue
		}
		if len(lg) < len(st) {
			lg = st
		} else if len(lg) == len(st) {
			lg = legiSmaller(lg, st)
		}
	}
	return lg
}

func legiSmaller(lg, st string) string {
	for i := range lg {
		if lg[i] == st[i] {
			continue
		} else if lg[i] < st[i] {
			return lg
		}
		return st
	}
	return lg
}

func findLongestHelper(s, st string) bool {
	m, n := len(s), len(st)
	tb := make([][]int, m+1)
	for i := range tb {
		tb[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				tb[i][j] = 0
			} else if s[i-1] == st[j-1] {
				tb[i][j] = tb[i-1][j-1] + 1
			} else {
				tb[i][j] = tb[i-1][j]
			}
		}
	}
	return tb[m][n] == n
}

var sum = 0

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	convertBST(root.Right)
	root.Val += sum
	sum = root.Val
	convertBST(root.Left)
	return root
}

func pathSum(root *TreeNode, sum int) [][]int {
	// dfs
	curArr := make([]int, 0)
	res := make([][]int, 0)
	pathSumHelper(root, sum, &curArr, &res)
	return res
}

func pathSumHelper(node *TreeNode, remain int, curArr *[]int, res *[][]int) {
	if node == nil {
		return
	}
	*curArr = append(*curArr, node.Val)
	if isLeaf(node) && node.Val == remain {
		t := make([]int, len(*curArr))
		copy(t, *curArr)
		*res = append(*res, t)
	}
	pathSumHelper(node.Left, remain-node.Val, curArr, res)
	pathSumHelper(node.Right, remain-node.Val, curArr, res)
	*curArr = (*curArr)[0 : len(*curArr)-1]
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
