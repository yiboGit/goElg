package algorithm

func qsort(arr []int) {
	qsortHelper(arr, 0, len(arr)-1)
}

func qsortHelper(arr []int, l, r int) {
	if l < r {
		h := partition(arr, l, r)
		qsortHelper(arr, l, h)
		qsortHelper(arr, h+1, r)
	}
}

func partition(arr []int, l, r int) int {
	pivot := arr[l]
	i, j := l, r
	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func TopK(arr []int, k int) int {
	return TopKHelper(arr, k, 0, len(arr)-1)
}

func TopKHelper(arr []int, k, l, r int) int {
	h := partition(arr, l, r)
	rank := h - l + 1
	if k == rank {
		return arr[h]
	}
	if k < rank {
		return TopKHelper(arr, rank-k, l, h)
	}
	return TopKHelper(arr, k-rank, h+1, r)
}
