// https://www.acwing.com/problem/content/description/1099/
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
	n, m int
	g    [N]string
	st   [N][N]bool
	q    [N * N]pair

	ans int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, N*N)
	Fscanf(reader, "%d %d\n", &n, &m)
	var s string
	for i := 0; i < n; i++ {
		Fscanf(reader, "%s\n", &s)
		g[i] = s
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !st[i][j] && g[i][j] == 'W' {
				dfs(i, j)
				ans++
			}
		}
	}
	Println(ans)

}

func dfs(x int, y int) {
	hh, tt := 0, -1
	tt++
	q[tt] = pair{x, y}
	for hh <= tt {
		head := q[hh]
		st[head.x][head.y] = true // 变更中心点的状态
		hh++
		// 八个方向拓展
		for xs := head.x - 1; xs <= head.x+1; xs++ {
			for ys := head.y - 1; ys <= head.y+1; ys++ {
				if xs == head.x && ys == head.y {
					continue
				}
				if xs < 0 || xs >= n || ys < 0 || ys >= m {
					continue
				}
				if st[xs][ys] || g[xs][ys] == '.' {
					continue
				}
				tt++
				q[tt] = pair{xs, ys}
				st[xs][ys] = true  // 变更周围点的状态
			}
		}
	}
}
