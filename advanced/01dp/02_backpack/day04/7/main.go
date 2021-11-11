/*
https://www.acwing.com/problem/content/7/
混合背包
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	f [N]int

	n, m int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var v, w, s int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &v, &w, &s)
		if s == 0 {
			// 完全背包
			for j := v; j <= m; j++ {
				f[j] = max(f[j], f[j-v]+w)
			}
		} else {
			if s == -1 {
				// 0,1 背包是特殊的多重背包
				s = 1
			}
			k := 1
			for k <= s {
				// 二进制优化多重背包
				for j := m; j >= k*v; j-- {
					f[j] = max(f[j], f[j-k*v]+k*w)
				}
				s -= k
				k *= 2
			}
			if s > 0 {
				for j := m; j >= s*v; j-- {
					f[j] = max(f[j], f[j-s*v]+s*w)
				}
			}
		}
	}
	fmt.Println(f[m])
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
