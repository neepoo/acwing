package main

import (
	"fmt"
	"os"
)

const N = 30

var (
	n, m int
	path [N]int
)

// u代表第k层, s代表第k层开始选的数
func dfs(u, s int) {
	if u > m {
		for i := 1; i <= m; i++ {
			fmt.Printf("%d ", path[i])

		}
		fmt.Println()
		return
	}
	for i := s; i <= n; i++ {
		path[u] = i
		dfs(u+1, i+1)
		path[u] = 0
	}
}
func main() {
	fmt.Fscanf(os.Stdin, "%d %d",&n, &m)
	dfs(1, 1)
}
