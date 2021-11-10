package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1010

type pair struct {
	x, y int
}

var (
	g  [N][N]int
	p  [N][N]pair
	q  [N * N]pair
	hh = 0
	tt = -1
	n  int

	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func bfs() {
	tt++
	q[tt] = pair{0, 0}
	for hh <= tt {
		head := q[hh]
		hh++
		hx, hy := head.x, head.y
		// 遍历四个方向
		for i := 0; i < 4; i++ {
			x := hx + dx[i]
			y := hy + dy[i]

			if (x >= 0) && (x < n) && (y >= 0) && (y < n) && (g[x][y] == 0) && (p[x][y] == pair{-1, -1}) {
				p[x][y] = head
				tt++
				q[tt] = pair{x, y}
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N*N)
	Fscanf(reader, "%d\n", &n)
	var x int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Fscanf(reader, "%d ", &x)
			g[i][j] = x
			p[i][j] = pair{-1, -1}
		}
	}
	bfs()
	st := [N * N]pair{}
	ix, iy := n-1, n-1
	tx, ty := ix, iy
	idx := 0
	for {
		idx++
		st[idx] = pair{ix, iy}
		ix = p[tx][ty].x
		iy = p[tx][ty].y
		if ix == 0 && iy == 0 {
			idx++
			st[idx] = pair{0, 0}
			break
		}
		tx = ix
		ty = iy
	}
	for idx > 0 {
		Printf("%d %d\n", st[idx].x, st[idx].y)
		idx--
	}
}
