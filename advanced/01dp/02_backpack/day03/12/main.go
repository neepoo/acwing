package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	f    [N][N]int
	v, w [N]int
	n, m int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &v[i], &w[i])
	}

	for i := n; i >= 1; i-- {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i+1][j]
			if j >= v[i] {
				f[i][j] = max(f[i][j], f[i+1][j-v[i]] + w[i])
			}
		}
	}
	j := m
	for i:=1;i<=n;i++{
		if j >= v[i] && f[i][j] == f[i+1][j-v[i]]+w[i]{
			fmt.Printf("%d ", i)
			j-=v[i]
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
