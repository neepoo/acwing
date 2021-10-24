//https://www.acwing.com/problem/content/1100/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 55

type pair struct {
	x, y int
}

var (
	n, m int
	g    [N][N]int
	st   [N][N]bool
	q    [N * N]pair

	cnt, area int

	dx = [4]int{0, -1, 0, 1}
	dy = [4]int{-1, 0, 1, 0}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, N*N)
	Fscanf(reader, "%d %d\n", &n, &m)
	var x int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				Fscanf(reader, "%d\n", &x)
			} else {
				Fscanf(reader, "%d ", &x)
			}
			g[i][j] = x
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !st[i][j] {
				currentArea := dfs(i, j)
				area = max(area, currentArea)
				cnt++
			}
		}
	}
	Println(cnt)
	Println(area)
}

func max(a int, i int) int {
	if a > i {
		return a
	}
	return i
}

func dfs(i int, j int) int {
	currentArea := 0
	hh, tt := 0, -1
	tt++
	q[tt] = pair{i, j}
	st[i][j] = true
	for hh <= tt {
		currentArea++
		head := q[hh]
		hh++
		x := head.x
		y := head.y
		for k := 0; k < 4; k++ {
			a := x + dx[k]
			b := y + dy[k]
			// 是否在合法范围
			if a < 0 || a >= n || b < 0 || b >= m {
				continue
			}
			// 该房间是否遍历过
			if st[a][b] {
				continue
			}
			// 该方向是否是墙
			if ((g[x][y]>>k)&1) == 1 {
				continue
			}
			// 该方向能够探索
			tt++
			q[tt] = pair{a, b}
			st[a][b] = true
		}
	}
	return currentArea
}
