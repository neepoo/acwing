package main

// https://www.acwing.com/problem/content/844/

import (
	"bytes"
	"fmt"
)

const N = 20

var (
	col, dg, udg [N]bool
	q            [N][N]string
	n            int
	bf           bytes.Buffer
)

// x表示第几层
func dfs(x int) {
	if x == n+1 {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				bf.WriteString(q[i][j])
			}
			bf.WriteString("\n")
		}
		bf.WriteString("\n")
		return
	}
	// 判断第x行第i列是否能够合法放置
	for i := 1; i <= n; i++ {
		// 可行性剪枝
		if !col[i] && !dg[x-i+n] && !udg[x+i] {
			q[x][i] = "Q"
			col[i], dg[x-i+n], udg[x+i] = true, true, true
			dfs(x + 1)
			col[i], dg[x-i+n], udg[x+i] = false, false, false
			q[x][i] = "."
		}
	}
}

func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			q[i][j] = "."
		}
	}
	dfs(1)
	fmt.Print(bf.String())
}
