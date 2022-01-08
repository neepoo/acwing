package main

import . "fmt"

const N = 10

var (
	n   int
	ans int
	g   [N]string
	dx  = []int{-1, 0, 1, 0}
	dy  = []int{0, 1, 0, -1}
	st  [N][N]bool
)

func dfs(x, y, l, r int) {
	if l == r {
		ans = max(ans, l+r)
		return
	}
	st[x][y] = true
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a < 0 || a >= n || b < 0 || b >= n || st[a][b] {
			continue
		}
		if g[x][y] == ')' && g[a][b] == '(' {
			continue
		}
		if g[a][b] == '(' {
			dfs(a, b, l+1, r)
		} else {
			dfs(a, b, l, r+1)
		}
	}
	st[x][y] = false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	Scan(&n)
	for i := 0; i < n; i++ {
		Scan(&g[i])
	}
	if g[0][0] == '(' {
		dfs(0, 0, 1, 0)
	}

	Println(ans)
}
