package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	n int
	g [N][N]int
	f [N][N]int
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func dp() {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 {
				// 第一行只能从左往右走
				f[i][j] = g[i][j] + f[i][j-1]
			} else if j == 1 {
				// 第一列只能从上往下走
				f[i][j] = g[i][j] + f[i-1][j]
			} else {
				f[i][j] = g[i][j] + min(f[i-1][j], f[i][j-1])
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 20010)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscanf(reader, "%d ", &g[i][j])
		}
	}
	dp()
	fmt.Println(f[n][n])
}
