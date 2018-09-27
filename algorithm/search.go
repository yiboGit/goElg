package algorithm

func BinarySearch(arr []int, a int) int {
	return searchBetween(arr, a, 0, len(arr)-1)
}
func searchBetween(arr []int, a, left, right int) int {
	for {
		if left > right {
			return -1
		}
		m := left + (right-left)/2
		if arr[m] == a {
			return m
		} else if arr[m] > a {
			right = m - 1
		} else {
			left = m + 1
		}
	}
}
func SearchInSortedRotateArray(arr []int, a int) int {
	// 2,3,-2,-1,0,1   3
	l := len(arr)
	if l == 0 {
		return -1
	}
	left, right := 0, l-1
	for {
		if left > right {
			return -1
		}
		m := left + (right-left)/2
		if arr[m] == a {
			return m
		}
		if arr[m] < arr[right] {
			if a > arr[m] && a <= arr[right] {
				left = m + 1
			} else {
				right = m - 1
			}
		} else {
			if a >= arr[left] && a < arr[m] {
				right = m - 1
			} else {
				left = m + 1
			}
		}
	}
}

func pow(a, n int) int {
	if n == 0 {
		return 1
	}
	b := pow(a, n/2)
	if n%2 == 0 {
		return b * b
	}
	return a * b * b
}

type Ma struct {
	a11, a12, a21, a22 int
}

func (m Ma) mul(n Ma) Ma {
	return Ma{
		a11: m.a11*n.a11 + m.a12*n.a21,
		a12: m.a11*n.a12 + n.a12*n.a22,
		a21: m.a21*n.a11 + m.a22*n.a21,
		a22: m.a21*n.a12 + m.a22*n.a22,
	}
}

var f = Ma{1, 1, 1, 0}

func fib(n int) int {
	r := fibHelper(n)
	return r.a21

}
func fibHelper(n int) Ma {
	if n == 1 {
		return f
	}
	b := fibHelper(n / 2)
	b2 := b.mul(b)
	if n%2 == 0 {
		return b2
	}
	return b2.mul(f)
}
func fib2(n int) int {
	a, b := 0, 1
	var r int
	for i := 2; i <= n; i++ {
		r = a + b
		a = b
		b = r
	}
	return r
}
