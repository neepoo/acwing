// 二维费用 01背包 裸题
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 1010
	M = 110
)

var (
	n, v, m int
	f       [M][M]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d\n", &n, &v, &m)

	for i := 1; i <= n; i++ {
		var a, b, c int // 体积,重量和价值
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		for j := v; j >= a; j-- {
			for k := m; k >= b; k-- {
				f[j][k] = max(f[j][k], f[j-a][k-b]+c)
			}
		}
	}

	fmt.Println(f[v][m])
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
