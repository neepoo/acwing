package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acwing.com/problem/content/846/
type pair struct {
	i, j int
}

const (
	N = 110
)

var (
	dx   = [4]int{-1, 0, 1, 0}
	dy   = [4]int{0, 1, 0, -1}
	n, m int
	g    [N][N]int
	d    [N][N]int
	q    [N * N]pair
	hh   = 0
	tt   = -1
)

func bfs() int {
	tt++
	q[tt] = pair{0, 0}
	d[0][0] = 0
	for hh <= tt {
		head := q[hh]
		hh++
		for i := 0; i < 4; i++ {
			x := head.i + dx[i]
			y := head.j + dy[i]
			if x >= 0 && x < n && y >= 0 && y < m && g[x][y] == 0 && d[x][y] == -1 {
				d[x][y] = d[head.i][head.j] + 1
				tt++
				q[tt] = pair{x, y}
			}
		}
	}

	return d[n-1][m-1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, N*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var x int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == (m - 1) {
				fmt.Fscanf(reader, "%d\n", &x)
			} else {
				fmt.Fscanf(reader, "%d ", &x)
			}
			g[i][j] = x
			d[i][j] = -1
		}
	}
	fmt.Println(bfs())
}
