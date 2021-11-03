package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 110
	INF = 0x3f3f3f
)

var (
	m, n  int
	g     [N][N]int
	level [N]int

	dist [N]int
	st   [N]bool
)

func init() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			g[i][j] = INF
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &m, &n)
	var p, l, x int
	var t, v int
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &p, &l, &x)
		level[i] = l
		g[0][i] = min(g[0][i], p)
		// 读取替代品
		for j := 0; j < x; j++ {
			fmt.Fscanf(reader, "%d %d\n", &t, &v)
			g[t][i] = min(g[t][i], v)
		}
	}
	var res = INF
	// 酋长的级别是最高的, +m的意思是枚举所有可能的级别区间
	for i := level[1] - m; i <= level[1]; i++ {
		res = min(dijkstra(i, i+m), res)
	}
	fmt.Println(res)
}

func dijkstra(left int, right int) int {
	for i := 0; i < N; i++ {
		dist[i] = INF
		st[i] = false
	}

	dist[0] = 0
	// n+1个物品(包括0号物品,不交换, 直接买的)
	for i := 0; i < n; i++ {
		t := -1
		for j := 0; j <= n; j++ {
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}
		st[t] = true
		// 松弛
		for k := 1; k <= n; k++ {
			// 需要满足等级限制
			if level[k] >= left && level[k] <= right {
				if dist[k] > dist[t]+g[t][k] {
					dist[k] = dist[t] + g[t][k]
				}
			}
		}
	}
	return dist[1]
}
