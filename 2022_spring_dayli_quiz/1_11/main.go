// https://www.acwing.com/problem/content/1980/
// 前缀最大值 后缀最小值
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

const (
	INF = int(1e9)
	N   = 100010
)

var (
	preMax [N]int // preMax[i] 表示排序过后的ps中,该线段及之前线段另一边的最大值
	sufMin [N]int
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	ps := make([]pair, n+1)
	// DUMMY NODE
	ps[0] = pair{-INF, -INF}
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &ps[i].x, &ps[i].y)
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].x < ps[j].x
	})
	preMax[0] = -INF
	sufMin[n+1] = INF
	for i := 1; i <= n; i++ {
		preMax[i] = max(preMax[i-1], ps[i].y)
	}
	for i := n; i >= 1; i-- {
		sufMin[i] = min(sufMin[i+1], ps[i].y)
	}
	var cnt int
	for i := 1; i <= n; i++ {
		y := ps[i].y
		if y > preMax[i-1] && y < sufMin[i+1] {
			cnt++
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
