package main

// https://www.acwing.com/problem/content/844/

import (
	"bytes"
	"fmt"
	"strconv"
)

const N = 8

var (
	st [N]bool
	q  [N]int
	n  int
	bf bytes.Buffer
)

// x表示第几层
func dfs(x int) {
	if x == n+1 {
		for i := 1; i <= n; i++ {
			bf.WriteString(strconv.Itoa(q[i]) + " ")
		}
		bf.WriteString("\n")
		return
	}
	for i := 1; i <= n; i++ {
		if !st[i] {
			q[x] = i
			st[i] = true
			dfs(x + 1)
			st[i] = false
		}
	}
}

func main() {
	fmt.Scanf("%d", &n)
	dfs(1)
	fmt.Print(bf.String())
}
