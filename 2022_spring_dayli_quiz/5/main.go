package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

const N = 100010

type pair struct {
	x, y int
}
type S []pair

func (s S) Len() int {
	return len(s)
}

func (s S) Less(i, j int) bool {
	return s[i].x < s[j].x
}

func (s S) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var (
	n  int
	h  [N]int
	q  [N]pair
	qq S
)

func unique() int {
	cnt := 1
	last := h[1]
	for i := 1; i <= n; i++ {
		if h[i] != last {
			cnt++
			h[cnt] = h[i]
			last = h[i]
		}
	}
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscan(reader, &n)
	for i := 1; i <= n; i++ {
		Fscan(reader, &h[i])
	}
	n = unique()
	for i := 1; i <= n; i++ {
		q[i] = pair{h[i], i}
	}
	qq = q[1 : n+1]

	sort.Sort(qq)

	res, cnt := 1, 1
	for i := 1; i <= n; i++ {
		k := q[i].y
		if h[k-1] < h[k] && h[k+1] < h[k] {
			cnt--
		} else if h[k-1] > h[k] && h[k+1] > h[k] {
			cnt++
		}
		if q[i].x != q[i+1].x {
			res = max(res, cnt)
		}
	}
	Println(res)
}
