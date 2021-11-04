package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 160

var (
	c, r   int
	g      [N]string
	ix, iy int
	wx, wy int
	dx     = [8]int{-2, -1, 1, 2, 2, 1, -1, -2}
	dy     = [8]int{1, 2, 2, 1, -1, -2, -2, -1}
	d      [N][N]int
)

type pair struct {
	x, y int
}

func bfs() {
	var q [N * N]pair
	hh, tt := 0, -1
	tt++
	q[tt] = pair{ix, iy}
	d[ix][iy] = 0
	flag := false
	for hh <= tt {
		head := q[hh]
		hh++

		hx, hy := head.x, head.y

		// 遍历八个方向，看是否能过去
		for i := 0; i < 8; i++ {
			x := hx + dx[i]
			y := hy + dy[i]

			if x >= 0 && x < r && y >= 0 && y < c && d[x][y] == -1 && g[x][y] != '*' {
				d[x][y] = d[hx][hy] + 1
				tt++
				q[tt] = pair{x, y}
				if x == wx && y == wy {
					flag = true
					break
				}
			}
		}
		if flag {
			return
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscanf(reader, "%d %d\n", &c, &r)
	var s string
	for i := 0; i < r; i++ {
		Fscanf(reader, "%s\n", &s)
		for idx, ch := range s {
			if ch == 'K' {
				ix = i
				iy = idx
			}
			if ch == 'H' {
				wx = i
				wy = idx
			}
			d[i][idx] = -1
		}
		g[i] = s
	}
	bfs()
	Println(d[wx][wy])
}
