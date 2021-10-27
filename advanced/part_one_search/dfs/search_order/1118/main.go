package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 10

var (
	t          int
	n, m, x, y int

	st [N][N]bool

	dx = []int{-2, -1, 1, 2, 2, 1, -1, -2}
	dy = []int{1, 2, 2, 1, -1, -2, -2, -1}

	ans int
)

func clear() {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			st[i][j] = false
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		fmt.Fscanf(reader, "%d %d %d %d\n", &n, &m, &x, &y)
		dfs(x, y, 1)
		fmt.Println(ans)
		ans = 0
	}
}

// 外部搜索
func dfs(x, y, cnt int) {
	if cnt == n*m {
		ans++
	}
	st[x][y] = true
	for i := 0; i < 8; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a < 0 || a >= n || b < 0 || b >= m {
			continue
		}
		if st[a][b] {
			continue
		}
		dfs(a, b, cnt+1)
	}

	st[x][y] = false
}
