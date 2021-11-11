/*
井字形
 */
package main

import (
	"fmt"
)

const (
	N = (1 << 10) + 10
)

var (
	state []int    // 存储每一行的合法的存储方案(把每个数看成二进制,1放置国王,0不放)
	head  [N][]int // head[a]表示a对应的放置状态可以转移到的集合状态

	n, k int               // 棋盘大小和国王数量
	f    [12][110][N]int64 // f[i][j][k]表示第i行放置了j个国王并且第i行摆放的是状态为k的所有集合
	cnt  [N]int            // 统计每个合法的放置方案中有几个国王i

)

func check(x int) bool {
	for i := 0; i < n; i++ {
		if (((x >> i) & 1) == 1) && (((x >> (i + 1)) & 1) == 1){
			// 相邻两个格子可以相互攻击
			return false
		}
	}
	// 该状态小的每一个国王都不会互相攻击
	return true
}

// 统计x所表示的二进制数,包含几个1(即放置了几个国王)
func count(x int) (res int) {
	for i := 0; i < n; i++ {
		res += (x >> i) & 1
	}
	return
}

// 预处理
// 计算出哪些是可行的放置方案,以及
// 哪些合法的方案能转移到另外的合法方案
func pre() {
	for i := 0; i < (1 << n); i++ {
		if check(i) {
			// 存储所有无法左右攻击的方案
			state = append(state, i)
			cnt[i] = count(i)
		}
	}

	// 枚举所有的合法方案,记录可行的转移方案
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state); j++ {
			a := state[i]
			b := state[j]
			if (a&b) == 0 /*判断上下是否能够攻击到*/ && check(a|b) /*判断两个斜对角是否能够攻击到*/ {
				head[i] = append(head[i], j) // 第j个合法方案可以转移到第i个合法方案
			}
		}
	}
}

func main() {
	fmt.Scan(&n, &k)
	// 预处理得到合法方案以及合法的转移关系
	pre()
	// dp
	f[0][0][0] = 1
	for i := 1; i <= n+1; i++ { // 第几行
		for j := 0; j <= k; j++ { // 目前使用的棋子数量
			for a := 0; a < len(state); a++ {
				for b := 0; b < len(head[a]); b++ {
					c := cnt[state[a]]  // 统计第i行的国王数量
					if j >= c{
						// 							 b 能够转移到a状态
						f[i][j][a] += f[i-1][j-c][head[a][b]]
					}
				}
			}
		}
	}
	fmt.Println(f[n+1][k][0])
}
