package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 12

var (
	g [N][N]int
	f [2 * N][N][N]int

	n int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dp() {
	for k := 2; k <= 2*n; k++ {
		for i1 := 1; i1 <= n; i1++ {
			for i2 := 1; i2 <= n; i2++ {
				j1, j2 := k-i1, k-i2
				if j1 >= 1 && j1 <= n && j2 >= 1 && j2 <= n {
					t := g[i1][j1]
					if i1 != i2 {
						// 两条路径走到的不是同一个点
						t += g[i2][j2]
					}
					f[k][i1][i2] = max(f[k][i1][i2], f[k-1][i1-1][i2-1]+t)
					f[k][i1][i2] = max(f[k][i1][i2], f[k-1][i1][i2-1]+t)
					f[k][i1][i2] = max(f[k][i1][i2], f[k-1][i1-1][i2]+t)
					f[k][i1][i2] = max(f[k][i1][i2], f[k-1][i1][i2]+t)
				}
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var a, b, c int
	for {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		if a == 0 && b == 0 && c == 0 {
			break
		}
		g[a][b] = c
	}
	dp()
	fmt.Println(f[2*n][n][n])
}
