package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 110

var (
	t int

	g    [N][N]int
	r, c int

	f [N][N]int
)

func clear() {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			f[i][j] = 0
		}
	}
}

func dp() int {
	for i := 1; i<=r;i++{
		for j:=1; j <= c;j++{
			f[i][j] = g[i][j] + max(f[i-1][j], f[i][j-1])
		}
	}
	return f[r][c]
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 1e6)
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		fmt.Fscanf(reader, "%d %d\n", &r, &c)
		for j := 1; j <= r; j++ {
			for k := 1; k <= c; k++ {
				fmt.Fscanf(reader, "%d ", &g[j][k])
			}
		}
		fmt.Println(dp())
		clear()
	}

}
