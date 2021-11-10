package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 110
	INF = 0x3f3f3f3f
)

var (
	d    [N][N]int
	n, m int
)

func floyd() {
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
}
func init() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			d[i][j] = INF
			if i == j {
				d[i][j] = 0
			}
		}
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		v := min(d[a][b], c)
		d[a][b] = v
		d[b][a] = v
	}
	floyd()
	var res int
	for i := 1; i <= n; i++ {
		if d[1][i] == INF {
			res = -1
			break
		}
		res = max(res, d[1][i])
	}
	fmt.Println(res)
}

func min(i int, c int) int {
	if i < c {
		return i
	}
	return c
}

func max(i int, c int) int {
	if i < c {
		return c
	}
	return i
}
