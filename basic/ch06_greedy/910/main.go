package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

//https://www.acwing.com/problem/content/910/
const N = 100010

type pair struct {
	x, y int
}

var (
	n   int
	q   [N]pair
	g   [N]int
	res = 1
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscan(reader, &n)
	for i := 0; i < n; i++ {
		Fscan(reader, &q[i].x, &q[i].y)
	}
	sort.Slice(q[:n], func(i, j int) bool {
		return q[i].x < q[j].x
	})
	maxR := q[0].y
	for i := 1; i < n; i++ {
		// 暴力
		if q[i].x <= maxR{
			maxR = min(maxR, q[i].y)
		}else {
			res++
			maxR = q[i].y
		}
	}
	Println(res)
}

/*
10
-5 1
-10 7
-8 4
-2 -1
-7 -6
0 7
-8 -7
3 9
-9 -4
-3 -3
*/
func min(r int, j int) int {
	if r > j {
		return j
	}
	return r
}
