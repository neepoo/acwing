package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 55

var (
	n int
	g [N][N]int

	st1, st2 [N][N]bool

	x1, y1, x2, y2 int

	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func cost(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

type pair struct {
	x, y int
}

func bfs(x, y int, st *[N][N]bool) {
	var q [N * N]pair
	hh, tt := 0, -1
	tt++
	q[tt] = pair{x, y}
	st[x][y] = true
	for hh <= tt {
		head := q[hh]
		hh++
		for i := 0; i < 4; i++ {
			a := head.x + dx[i]
			b := head.y + dy[i]
			if a < 1 || a > n || b < 1 || b > n {
				continue
			}
			if g[a][b] == 1 {
				continue
			}
			if st[a][b] {
				continue
			}
			tt++
			q[tt] = pair{a, b}
			st[a][b] = true
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
	fmt.Fscanf(reader, "%d\n", &n)
	fmt.Fscanf(reader, "%d %d\n", &x1, &y1)
	fmt.Fscanf(reader, "%d %d\n", &x2, &y2)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j == n {
				fmt.Fscanf(reader, "%1d\n", &g[i][j])
			} else {
				fmt.Fscanf(reader, "%1d", &g[i][j])
			}
		}
	}
	var ans = 0x3f3f3f3f

	bfs(x1, y1, &st1)
	if st1[x2][y2] {
		fmt.Println(0)
		return
	} else {
		bfs(x2, y2, &st2)
		for i1 := 1; i1 <= n; i1++ {
			for j1 := 1; j1 <= n; j1++ {
				if st1[i1][j1] {
					for i2 := 1; i2 <= n; i2++ {
						for j2 := 1; j2 <= n; j2++ {
							if st2[i2][j2] {
								ans = min(ans, cost(i1, j1, i2, j2))
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
