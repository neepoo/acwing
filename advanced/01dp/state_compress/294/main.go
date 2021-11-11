package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 110
	M = 11
	S = 1 << M
)

var (
	n, m int

	// 已经摆放的炮兵不会相互攻击到,并且每个炮兵都在平地
	// f[i][j][k] 表示第i行之前已经摆放好,且第i行的状态是k,第i-1行的状态j的摆放方案的最大值并且
	f [2][S][S]int

	g [N]int // 用来表示地形图

	state []int // 熟悉的朋友,用来存储所有合法的方案

	cnt []int // 用来表示当前合法方案放置了几个炮

)

func count(x int) (res int) {
	for x > 0 {
		res += x & 1
		x >>= 1
	}
	return
}

// 用来判断当前行的状态是否合法, 当前位为炮的话,相邻的两位都必须不是炮.注意这里不需要考虑高地的情况
func check(x int) bool {
	return x >> 1 & x == 0 && x >> 2 & x == 0
}

//func check(x int) bool {
//	for i := 0; i < m; i++ {
//		if x>>i+1 == 1 || x>>i+2 == 1 {
//			return false
//		}
//	}
//	return true
//}

// 预处理
func pre() {
	// 得到所有的合法方案及其相对应的炮数量
	for i := 0; i < (1 << m); i++ {
		if check(i) {
			state = append(state, i)
			cnt = append(cnt, count(i))
		}
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var x string
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%s\n", &x)
		for j := 0; j < len(x); j++ {
			if x[j] == 'H' {
				g[i] += 1 << j
			}
		}
	}
	// 预处理
	pre()
	//dp
	for i := 1; i <= n+2; i++ {
		for j, b := range state { // 第i-1行的摆放
			for k, c := range state { // 第i行的摆放
				for u, a := range state { // 第i-2行的摆放
					if a&b != 0 || a&c != 0 || b&c != 0 {
						//i-2到i行, 某一列上的炮冲突了
						continue
					}
					if g[i]&c != 0 {
						//第i行有高地, 摆放不了
						continue
					}
					f[i&1][j][k] = max(f[i&1][j][k], f[(i-1)&1][u][j]+cnt[k])

				}
			}
		}
	}
	fmt.Println(f[(n+2)&1][0][0])
}
