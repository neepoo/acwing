package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 2010
)

var (
	n, m int
	g    [N][N]float64
	dist [N]float64
	s, t int
	st   [N]bool
)

func init() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			g[i][j] = 1.0
		}
	}
}

func dijkstra() {
	for i := 0; i <= n; i++ {
		dist[i] = float64(0x3f3f3f3f)
	}
	// 从终点反向搜索
	dist[t] = 100.0
	for i := 0; i < n -1; i++ {
		// 找到距离终点转账开销最小的点
		k := -1
		for j := 1; j <= n; j++ {
			if !st[j] && (k == -1 || dist[j] < dist[k]) {
				k = j
			}
		}
		st[k] = true
		for j:=1;j<=n;j++{
			if g[k][j] > 0 && g[k][j] < 1{
				// 说明这两者可以转账
				dist[j] = min(dist[j], dist[k] / (1-g[k][j]))
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 400000)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		g[a][b] = min(g[a][b], float64(c)/100.0)
		g[b][a] = g[a][b]
	}
	fmt.Fscanf(reader, "%d %d\n", &s, &t)
	dijkstra()
	fmt.Printf("%0.8f", dist[s])
}

func min(f float64, i float64) float64 {
	if f < i {
		return f
	}
	return i
}
