package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const N = 1010

var (
	g, f [N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t, n, res int
	fmt.Fscanf(reader, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(reader, "%d\n", &n)
		for j := 1; j <= n; j++ {
			if j < n {
				fmt.Fscanf(reader, "%d ", &g[j])
			} else {
				fmt.Fscanf(reader, "%d\n", &g[j])
			}
		}

		//正向
		for i := 1; i <= n; i++ {
			f[i] = 1
			for j := 1; j < i; j++ {
				if g[i] > g[j] {
					f[i] = max(f[i], f[j]+1)
					res = max(res, f[i])
				}
			}
		}

		// 反向
		for i := 1; i <= n; i++ {
			f[i] = 1
			for j := 1; j < i; j++ {
				if g[i] < g[j] {
					f[i] = max(f[i], f[j]+1)
					res = max(res, f[i])
				}
			}
		}
		fmt.Println(res)
		res = 0
	}
}
