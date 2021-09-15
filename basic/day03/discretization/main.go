package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

//https://www.acwing.com/problem/content/804/
const N = 100010

type Pair struct {
	first  int
	second int
}

func NewPair(first int, second int) Pair {
	return Pair{first: first, second: second}
}

var (
	n, m  int
	x, c  int
	l, r  int
	all   []int
	add   []Pair
	query []Pair
	a     [3 * N]int
	s     [3 * N]int
)

func unique(a []int) []int {
	i := 0
	for j := 0; j < len(a); j++ {
		if j == 0 || a[j] != a[j-1] {
			a[i] = a[j]
			i++
		}
	}
	return a[:i]
}

func find(x int) int {
	l := 0
	r := len(all) - 1
	for l < r {
		mid := (l + r) / 2
		if all[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r + 1
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 10*N)
	Fscanf(reader, "%d %d\n", &n, &m)
	for i := 0; i < n; i++ {
		Fscanf(reader, "%d %d\n", &x, &c)
		add = append(add, NewPair(x, c))
		all = append(all, x)
	}
	for i := 0; i < m; i++ {
		Fscanf(reader, "%d %d\n", &l, &r)
		query = append(query, NewPair(l, r))
		all = append(all, l, r)
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i] < all[j]
	})
	all = unique(all)

	for i := 0; i < n; i++ {
		x = add[i].first
		c = add[i].second
		idx := find(x)
		a[idx] += c
	}
	// 处理前缀和
	for i := 1; i <= len(all); i++ {
		s[i] += a[i] + s[i-1]
	}
	for i := 0; i < m; i++ {
		l = query[i].first
		r = query[i].second
		ii := find(l)
		ir := find(r)
		Println(s[ir] - s[ii-1])
	}
}
