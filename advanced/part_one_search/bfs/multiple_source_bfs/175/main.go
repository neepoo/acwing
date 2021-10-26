package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

const N = 1010

var (
	n, m int
	g    [N][N]int
	dist [N][N]int

	q  [N * N]pair
	hh = 0
	tt = -1

	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func _init() {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dist[i][j] = -1
		}
	}
}

func bfs() {
	for hh <= tt {
		head := q[hh]
		hh++
		hx := head.x
		hy := head.y

		for i := 0; i < 4; i++ {
			x := hx + dx[i]
			y := hy + dy[i]
			if x < 0 || x >= n || y < 0 || y >= m {
				continue
			}
			if dist[x][y] != -1 {
				continue
			}
			dist[x][y] = dist[hx][hy] + 1
			tt++
			q[tt] = pair{x, y}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	_init()
	var x int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Fscanf(reader, "%1d\n", &x)
			} else {
				fmt.Fscanf(reader, "%1d", &x)
			}
			g[i][j] = x
			if x == 1 {
				// 入队
				tt++
				q[tt] = pair{i, j}
				dist[i][j] = 0
			}
		}
	}
	bfs()
	var buf = &bytes.Buffer{}
	bufio.NewWriterSize(buf, 2*N*N)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprintf(buf, "%d ", dist[i][j])
		}
		fmt.Fprintf(buf, "%s", "\n")
	}
	fmt.Print(buf.String())
}
