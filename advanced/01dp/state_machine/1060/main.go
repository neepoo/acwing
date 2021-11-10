package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 100010
	INF = 0x3f3f3f3f
)

var (
	n int
	w [N]int
	/*
		f[i][0]  表示第i天持有股票的最大收益
		f[i][1]  表示第i-1天卖出股票,第i天属于冷冻期的最大收益
		f[i][2]  表示第i天依旧观望,没有买入股票,也没有卖出股票
	*/
	f [N][3]int
)

func init() {
	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			f[i][j] = -INF
		}
	}
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &w[i])
	}
	f[0][2] = 0
	for i := 1; i <= n; i++ {
		f[i][0] = max(f[i-1][2] - w[i], f[i-1][0])
		f[i][1] = f[i-1][0] + w[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	fmt.Println(max(f[n][1], f[n][2]))
}
