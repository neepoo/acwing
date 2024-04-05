package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 210
	MAX = 0x3f3f3f3f
)

var (
	n, m, k int
	g       [N][N]int
)

func init() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			g[i][j] = MAX
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func floyd() {
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m, &k)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		g[a][b] = min(g[a][b], c)
	}
	floyd()
	for i := 0; i < k; i++ {
		fmt.Fscan(r, &a, &b)
		if g[a][b] >= MAX/2 {
			fmt.Println("impossible")
		} else {
			fmt.Println(g[a][b])
		}
	}
}
