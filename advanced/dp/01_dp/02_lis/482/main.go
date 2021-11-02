package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	a, f, g [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n, ans int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &a[i])
	}

	for i := 1; i <= n; i++ {
		f[i] = 1
		for j := 1; j < i; j++ {
			if a[i] > a[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}

	for i := n; i >= 1; i-- {
		g[i] = 1
		for j := n; j > i; j-- {
			if a[i] > a[j] {
				g[i] = max(g[i], g[j]+1)
			}
		}
	}

	for i := 1; i <= n; i++ {
		ans = max(f[i]+g[i]-1, ans)
	}
	fmt.Println(n - ans)
}
