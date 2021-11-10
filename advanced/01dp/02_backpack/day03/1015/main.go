package main

/*
 step1: 求解最大利润
	1.1 将每个公司看成物品组
	1.2 分到的物品数量当成体积
	1.3 所分到i件物品所产生的盈利当成价值
	1.4 问题转换成,在总体积不超过m的情况下, 从n个物品组中选,每个物品组只能选一个物品所能得到的最大收益是多少
*/

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 12
	M = 17
)

var (
	f [N][M]int

	n, m int
	w    [N][M]int
	way  [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscanf(reader, "%d ", &w[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			for k := 1; k <= m; k++ {
				if j >= k {
					f[i][j] = max(f[i][j], f[i-1][j-k]+w[i][k])
				}
			}
		}
	}
	fmt.Println(f[n][m])
	// 从后往前看, 第i家公司分了几件产品
	j := m
	for i := n; i >= 1; i-- {
		for k := 0; k <= j; k++ {
			if f[i][j] == f[i-1][j-k]+w[i][k] {
				way[i] = k
				j -= k
				break
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("%d %d\n", i, way[i])
	}

}
