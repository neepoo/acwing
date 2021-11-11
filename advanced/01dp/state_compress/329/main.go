/*
十字型

当前行的摆放只与上一行有关系
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 14
	M   = 1<<12 + 10
	mod = 1e8
)

var (
	n, m int
	g    [N]int
	f    [N][M]int

	state []int    // 无法知道事先有多少种合法的方案
	head  [M][]int // 记录能够转移的合法方案
)

func check(x int) bool {
	for i := 0; i+1 < m; i++ {
		// 只要相邻的两块地是玉米就不合法
		if ((x>>i)&1) == 1 && ((x>>(i+1))&1) == 1 {
			return false
		}
	}
	return true
}

func pre() {
	// 预处理得到所有的合法方案
	for i := 0; i < (1 << m); i++ {
		if check(i) {
			state = append(state, i)
		}
	}

	// 处理得到合法状态转移的head矩阵
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state); j++ {
			// 得到方案
			a := state[i]
			b := state[j]
			if a&b == 0 {
				// 第j个方案可以转移到第i个方案
				head[i] = append(head[i], j)
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var x int
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d ", &x)
			// 0 -> 1 方便到时候&运算,表示不让种玉米

			if x == 0 {
				g[i] += 1 << j
			}
		}
	}

	// 预处理
	pre()
	// dp
	f[0][0] = 1
	// 枚举第i行
	for i := 1; i <= n+1; i++ {
		// 枚举第j个合法方案
		for j := 0; j < len(state); j++ {
			// 判断能不能种玉米
			if state[j]&g[i] == 0 {
				for _, k := range head[j] {
					f[i][j] = (f[i][j] + f[i-1][k]) % mod
				}
			}
		}
	}
	var res int
	for i := 0; i < (1 << m); i++ {
		res = (res + f[n][i]) % mod
	}
	fmt.Println(res)
	//fmt.Println(f[n+1][0])
}
