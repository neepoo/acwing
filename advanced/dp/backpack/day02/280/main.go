// 01背包求方案数量

package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 10010

var (
	n, m int

	f [110][M]int
)

func main() {
	f[0][0] = 1
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var x int
	for i := 1; i <= n; i++ { // 第一维枚举物品
		fmt.Fscanf(reader, "%d ", &x)
		for j := 0; j <= m; j++ { // 枚举体积
			f[i][j] = f[i-1][j]
			if j >= x {
				f[i][j] += f[i-1][j-x]
			}
		}
	}
	fmt.Println(f[n][m])
}
