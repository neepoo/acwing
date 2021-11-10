package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1010

var (
	n            int
	g            [N][N]int
	st           [N][N]bool
	// 队列开到bfs函数中会超时,以后要吸取教训!!
	q            = [N * N]pair{}
	peak, valley int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N*N)
	Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var x int
			if j == n-1 {
				Fscanf(reader, "%d\n", &x)
			} else {
				Fscanf(reader, "%d ", &x)
			}
			g[i][j] = x
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !st[i][j] {
				hasHeight, hasLower := bfs(i, j)
				// 周围没有比他高的,他就是山峰
				if !hasHeight {
					peak++
				}
				// 周围没有比他矮的,他就是山谷
				if !hasLower {
					valley++
				}
			}
		}
	}
	Println(peak, valley)
}

type pair struct {
	x, y int
}

func bfs(sx int, sy int) (bool, bool) {
	hasHeight, hasLower := false, false
	hh, tt := 0, -1

	st[sx][sy] = true
	tt++
	q[tt] = pair{sx, sy}
	for hh <= tt {
		head := q[hh]
		hh++

		// 八连通
		for i := head.x - 1; i <= head.x+1; i++ {
			for j := head.y - 1; j <= head.y+1; j++ {
				if i == head.x && j == head.y {
					continue
				}
				if (i >= 0) && (i < n) && (j >= 0) && (j < n) {
					if g[i][j] != g[head.x][head.y] {
						if g[i][j] > g[head.x][head.y] {
							hasHeight = true
						} else {
							hasLower = true
						}
						// 属于该连通块,并且没有访问过
					} else if !st[i][j] {
						st[i][j] = true
						tt++
						q[tt] = pair{i, j}
					}
				}
			}
		}
	}
	return hasHeight, hasLower
}
