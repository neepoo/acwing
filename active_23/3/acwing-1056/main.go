package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 5

var (
	n int
	a [N]int
	f [N]int // f[i]表示在前i天最多完成一笔交易的最大值
	g [N]int // g[i]表示在第i天后最多完成一笔交易的最大值
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<16)
	fmt.Fscan(reader, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var (
		minV = a[1]
		maxV = a[n]
		res  int
	)
	for i := 2; i <= n; i++ {
		f[i] = max( /*第i天不交易*/ f[i-1] /*第i天参与交易(卖出)*/, a[i]-minV)
		minV = min(minV, a[i])
	}
	for j := n - 1; j >= 1; j-- {
		g[j] = max( /*第j天参与卖出*/ maxV-a[j] /*第j天没参与卖出*/, g[j+1])
		maxV = max(maxV, a[j])
	}
	for k := 1; k <= n; k++ {
		res = max(res, f[k]+g[k+1])
	}
	fmt.Println(res)
}
