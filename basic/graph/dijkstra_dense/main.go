package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N   = 510
	MAX = 0x3f3f3f3f3f
)

var (
	n, m int
	g    [N][N]int
	dist [N]int
	st   [N]bool
)

func init() {
	for i := 0; i < N; i++ {
		dist[i] = MAX
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			g[i][j] = MAX
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dijkstraDense() int {
	dist[1] = 0
	for i := 0; i < n-1; i++ {
		t := -1
		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[j] < dist[t]) {
				t = j
			}
		}
		st[t] = true
		for k := 1; k <= n; k++ {
			if dist[k] > dist[t]+g[t][k] {
				dist[k] = dist[t] + g[t][k]
			}
		}
	}
	if dist[n] == MAX {
		return -1
	}
	return dist[n]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	Fscan(r, &n, &m)
	var x, y, z int
	for i := 0; i < m; i++ {
		Fscan(r, &x, &y, &z)
		g[x][y] = min(g[x][y], z)
	}
	Println(dijkstraDense())
}
