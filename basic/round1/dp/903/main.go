package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 310

var (
	r, c int
	g    [N][N]int
	f    [N][N]int

	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dp(x, y int) int {
	if f[x][y] != -1 {
		return f[x][y]
	}
	f[x][y] = 1
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a >= 1 && a <= r && b >= 1 && b <= c && g[x][y] > g[a][b] {
			f[x][y] = max(f[x][y], dp(a, b)+1)
		}
	}
	return f[x][y]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &r, &c)
	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			fmt.Fscanf(reader, "%d ", &g[i][j])
			f[i][j] = -1 // 还没访问过
		}
	}
	var res int

	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			res = max(res, dp(i, j))
		}
	}
	fmt.Println(res)
}
