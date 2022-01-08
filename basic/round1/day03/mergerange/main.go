package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

//https://www.acwing.com/problem/content/805/
const (
	N   = 100010
	INF = int(2e9)
)

var (
	n    int
	l, r int
	rs   []Pair

	res []Pair

	st = -INF
	ed = -INF
)

type Pair struct {
	first  int
	second int
}

func NewPair(first int, second int) Pair {
	return Pair{first: first, second: second}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 5*N)
	Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		Fscanf(reader, "%d %d\n", &l, &r)
		rs = append(rs, NewPair(l, r))
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].first < rs[j].first
	})
	for _, span := range rs {
		if span.first > ed {
			if st != -INF {
				res = append(res, NewPair(st, ed))
			}
			st = span.first
			ed = span.second
		} else {
			ed = max(ed, span.second)
		}
	}
	if st != -INF {
		res = append(res, NewPair(st, ed))
	}
	Println(len(res))
}
