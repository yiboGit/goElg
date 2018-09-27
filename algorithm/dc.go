package algorithm

import (
	"fmt"
	"math"
)

// 子数组的最大和 Dp  sum(i) = max(sum(i-1)+a[i], a[i])
func MaxSumOfSubArray(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	sum := make([]int, len(arr))
	sum[0] = arr[0]
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		add := sum[i-1] + arr[i]
		if add > arr[i] {
			sum[i] = add
		} else {
			sum[i] = arr[i]
		}
		if sum[i] > max {
			max = sum[i]
		}
	}
	return max
}

func CutRod(p []int, n int) int {
	r := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		t := math.MinInt64
		for j := 0; j < i; j++ {
			cp := p[j]
			pr := cp + r[i-j-1]
			t = int(math.Max(float64(t), float64(pr)))
		}
		r[i] = t
	}
	return r[n]
}

func LongestPandir(str string) string {
	type r struct {
		r int
		l int
		j int
	}
	result := r{0, 0, 0}
	l := len(str)
	i := 0
	for i < l {
		left, right := i, i
		for left > 0 && str[left] == str[i] {
			left--
		}
		left++
		for right < l && str[right] == str[i] {
			right++
		}
		right--
		j := 1
		for ; j < l>>1+1; j++ {
			if right+j >= l || j > left {
				break
			}
			if str[right+j] != str[left-j] {
				break
			}
		}
		j--
		if j > result.j {
			result.l = left - j
			result.r = right + j
			result.j = j
		}
		i++
	}
	return str[result.l : result.r+1]
}

func findMin(arr []int) int {
	return findHelper(arr, 0, len(arr)-1)
}

func findHelper(arr []int, l, r int) int {
	if l < r {
		m := l + (r-l)/2
		lv := findHelper(arr, l, m)
		rv := findHelper(arr, m+1, r)
		if lv < rv {
			return lv
		}
		return rv
	}
	return arr[l]
}

/**
 * Definition for a binary tree node.

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BST kth largest val
func kthSmallest(root *TreeNode, k int) int {
	cur := 1
	res := 0
	Tranverse(root, k, &cur, &res)
	return res
}

func Tranverse(node *TreeNode, k int, cur, res *int) {
	if node.Left != nil {
		Tranverse(node.Left, k, cur, res)
	}
	// Visit(node.Val, k, cur, res)
	if *cur == k {
		*res = node.Val
		return
	} else {
		*cur++
	}
	if node.Right != nil {
		Tranverse(node.Right, k, cur, res)
	}
}

// 递归构建binary tree
func BuildSearchTree(arr []int, l, r int) *TreeNode {
	if l <= r {
		m := l + (r-l)/2
		nd := &TreeNode{Val: arr[m]}
		nd.Left = BuildSearchTree(arr, l, m-1)
		nd.Right = BuildSearchTree(arr, m+1, r)
		return nd
	}
	return nil
}

func printTree(n *SegmentTreeNode) {
	if n == nil {
		return
	}
	fmt.Println(n.Val)
	printTree(n.Left)
	printTree(n.Right)
}

func RangeSum(arr []int, l, r int) {
}

// segment tree

type SegmentTreeNode struct {
	l, r  int
	Left  *SegmentTreeNode
	Right *SegmentTreeNode
	Val   int
}
type SegmentTree struct {
	root *SegmentTreeNode
}

// 范围的和
func BuildSegmentTree(arr []int, l, r int) *SegmentTreeNode {
	if l == r {
		nd := &SegmentTreeNode{l: l, r: r, Val: arr[l]}
		return nd
	}
	m := l + (r-l)>>1
	nd := &SegmentTreeNode{l: l, r: r, Val: 0}
	nd.Left = BuildSegmentTree(arr, l, m)
	if nd.Left != nil {
		nd.Val += nd.Left.Val
	}
	nd.Right = BuildSegmentTree(arr, m+1, r)
	if nd.Right != nil {
		nd.Val += nd.Right.Val
	}
	return nd
}

func QueryRange(arr []int, node *SegmentTreeNode, ql, qr int) int {
	if ql <= node.l && qr >= node.r {
		return node.Val
	}
	lv := 0
	if node.Left != nil {
		lv = QueryRange(arr, node.Left, ql, qr)
	}
	rv := 0
	if node.Right != nil {
		rv = QueryRange(arr, node.Right, ql, qr)
	}
	return lv + rv
}

func preorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	preHelper(root, &arr)
	return arr
}
func preHelper(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	preHelper(node.Left, res)
	preHelper(node.Right, res)
}

// 0 -> lo  0
// lo -> hi 1
// hi -> end 2
func sortColors(nums []int) {
	if len(nums) == 1 {
		return
	}
	c1, c2, c3 := 0, 0, len(nums)-1
	for c1 <= c3 {
		if nums[c1] == 0 {
			swap(&nums[c1], &nums[c2])
			c1++
			c2++
		} else if nums[c1] == 2 {
			swap(&nums[c1], &nums[c3])
			c3--
		} else {
			c1++
		}
	}
}
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func findDuplicates(arr []int) {
	for _, v := range arr {
		index := abs(v)
		if arr[index] >= 0 {
			arr[index] = arr[index] * -1
		} else {
			fmt.Print(index)
		}
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return x * -1
}

func findDuplicateHalf(arr []int) int {
	l := len(arr)
	if l == 1 {
		return arr[0]
	}
	c := arr[0]
	time := 1
	for i := 1; i < l; i++ {
		if time == 0 {
			time = 1
			c = arr[i]
			continue
		} else if arr[i] == c {
			time++
		} else {
			time--
		}
	}
	return c
}

func RotatePrintMatrix(arr [][]int) {
	i, j, m, n := 0, 0, len(arr), len(arr[0])
	for i < m && j < n {
		printHelper(arr, i, j, m, n)
		i++
		j++
		m--
		n--
	}
}
func printHelper(arr [][]int, i, j, m, n int) {
	i0, j0 := i, j
	for j0 < n {
		fmt.Print(arr[i0][j0])
		j0++
	}
	j0--
	i0++
	for i0 < m {
		fmt.Print(arr[i0][j0])
		i0++
	}
	i0--
	j0--
	for j0 >= j {
		fmt.Print(arr[i0][j0])
		j0--
	}
	j0++
	for i0 > i {
		fmt.Print(arr[i0][j0])
		i0--
	}
}
