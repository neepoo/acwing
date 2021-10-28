package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 5010

type pair struct {
	x, y int
}

func newPair(x int, y int) pair {
	return pair{x: x, y: y}
}

var (
	n   int
	ps  [N]pair
	f   [N]int
	ans int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var x, y int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &x, &y)
		ps[i] = newPair(x, y)
	}
	sort.Slice(ps[:n], func(i, j int) bool {
		return ps[i].x < ps[j].x
	})
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if ps[i].y > ps[j].y {
				f[i] = max(f[i], f[j]+1)
				ans = max(ans, f[i])
			}
		}
	}
	fmt.Println(ans)
}
