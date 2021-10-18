//https://www.acwing.com/problem/content/2/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	n, m int
	v, w [N]int
	f2   [N][N]int

	f1 [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func twoDimension() int {
	for i := 1; i <= n; i++ { // 枚举前i件物品
		for j := 0; j <= m; j++ {
			f2[i][j] = f2[i-1][j]
			if j >= v[i] {
				f2[i][j] = max(f2[i][j], f2[i-1][j-v[i]]+w[i])
			}
		}
	}
	return f2[n][m]
}

func oneDimension() int {
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			f1[j] = max(f1[j] /*第i个物品不要*/, f1[j-v[i]]+w[i])
		}
	}

	return f1[m]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 5*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 1; i <= n; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		v[i] = a
		w[i] = b
	}
	fmt.Println(oneDimension())
}
