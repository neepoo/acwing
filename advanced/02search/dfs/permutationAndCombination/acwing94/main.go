package main

import (
	"fmt"
)

const N = 20

var (
	n  int
	st [N]bool
)

func main() {
	fmt.Scanf("%d\n", &n)
	dfs(1)
}

// dfs 的u一般表示第几层
func dfs(u int) {
	if u > n {
		for i := 1; i <= n; i++ {
			if st[i] {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
		return
	}

	st[u] = true
	dfs(u + 1)
	st[u] = false
	dfs(u + 1)

}
