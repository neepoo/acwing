// 完全背包求方案数量
package main

import "fmt"

const N = 1010

var (
	m = []int{0, 10, 20, 50, 100}

	f [5][N]int
	n int
)

func main() {
	fmt.Scanf("%d", &n)
	f[0][0] = 1

	for i := 1; i <= 4; i++ { // 枚举物品
		// 枚举体积
		for j := 0; j <= n; j++ {
			// 枚举策略
			for k := 0; k*m[i] <= j; k++ {
				// 选0个第i个物品,1个,...,n个
				f[i][j] += f[i-1][j-k*m[i]]
			}
		}
	}

	fmt.Println(f[4][n])

}
