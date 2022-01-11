// https://www.acwing.com/problem/content/805/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	INF = int(2e9)
)

type pair struct {
	x, y int
}
type ps []pair

func (p ps) Len() int {
	return len(p)
}

func (p ps) Less(i, j int) bool {
	return p[i].x < p[j].x
}

func (p ps) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var (
	n    int
	segs ps
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	segs = make(ps, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &segs[i].x, &segs[i].y)
	}
	sort.Sort(segs)
	var (
		cnt = 0
		et  = -INF
	)
	for _, seg := range segs {
		if seg.x > et {
			if et == -INF {
				// 第一个区间
				cnt = 1
				et = seg.y
				continue
			}
			cnt++
			et = seg.y
		} else {
			et = max(et, seg.y)
		}
	}

	fmt.Println(cnt)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
