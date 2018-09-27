package algorithm

import (
	"fmt"
	"math/rand"
	"time"
)

type a interface {
	len()
}

func Shuffle(a []int) {
	l := len(a)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range a {
		next := i + rand.Intn(l-i)
		a[i], a[next] = a[next], a[i]
	}
}

func Brackets(n int) {
	BracketsHelper("", 0, 0, n)
}
func BracketsHelper(str string, open, close, n int) {
	if open == n && close == n {
		fmt.Println(str)
	}
	if open < n {
		BracketsHelper(str+"{", open+1, close, n)
	}
	if close < open {
		BracketsHelper(str+"}", open, close+1, n)
	}
}

func Permute(s string) {
	PermuteHelper("", s)
}

func PermuteHelper(h, s string) {
	if s == "" {
		fmt.Println(h)
	}
	t := []rune(s)
	for i := range t {
		t[i], t[0] = t[0], t[i]
		PermuteHelper(h+string(t[0]), string(t[1:]))
		t[i], t[0] = t[0], t[i]
	}
}

func d() {
	a := make([]string, 0)
	e(&a)
}

func e(arr *[]string) {
	*arr = append(*arr, "hah")
}
