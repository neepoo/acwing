package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 110

var (
	k int
	n int

	g [N]string

	xa, ya, xb, yb int

	st [N][N]bool

	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 1000000)
	fmt.Fscanf(reader, "%d\n", &k)
	for i := 0; i < k; i++ {
		fmt.Fscanf(reader, "%d\n", &n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%s\n", &g[j])
		}
		fmt.Fscanf(reader, "%d %d %d %d\n", &xa, &ya, &xb, &yb)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				st[i][j] = false
			}
		}
		if dfs(xa, ya) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func dfs(x, y int) bool {
	if g[x][y] == '#' {
		return false
	}
	if x == xb && y == yb {
		return true
	}
	st[x][y] = true
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a < 0 || a >= n || b < 0 || b >= n {
			continue
		}
		if st[a][b] {
			continue
		}
		if dfs(a, b) {
			return true
		}
	}
	return false
}
