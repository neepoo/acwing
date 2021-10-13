package main

import "fmt"

const N = 8

var (
	st [N]bool
	q  [N]int
	n  int
)

// x表示第几层
func dfs(x int) {
	if x == n+1 {
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", q[i])
		}
		fmt.Println()
		return
	}
	for i := 1; i <= n ; i++ {
		if !st[i]{
			q[x] = i
			st[i] = true
			dfs(x+1)
			st[i] = false
		}
	}
}

func main() {
	fmt.Scanf("%d", &n)
	dfs(1)
}
