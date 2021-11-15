/*
https://www.acwing.com/problem/content/1070/
介绍了如何将环展开成链
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 410
	INF = 0x3f3f3f3f
)

var (
	n    int
	w, s [N]int
	// f->min, g->max
	f, g [N][N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = INF
			g[i][j] = -INF
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)

	// 将环展开成链
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &w[i])
		w[i+n] = w[i]
	}

	// 处理前缀和
	for i := 1; i <= 2*n; i++ {
		s[i] = s[i-1] + w[i]
	}

	// dp
	// 先枚举区间长度
	for length := 1; length <= n; length++ {
		// 枚举区间的左端点
		for l := 1; l+length-1 <= 2*n; l++ {
			// 获取区间右端点坐标
			r := l + length - 1
			// 如果区间长度是一,那么代价是0
			if length == 1 {
				f[l][r], g[l][r] = 0, 0
			}
			// k=1表示最左边的一堆,只有一个石子
			for k := l ; k < r; k++ {
				f[l][r] = min(f[l][r], f[l][k]+f[k+1][r]+s[r]-s[l-1])
				g[l][r] = max(g[l][r], g[l][k]+g[k+1][r]+s[r]-s[l-1])
			}
		}
	}

	// 获取答案
	maxV, minV := -INF, INF
	for i := 1; i <= n; i++ {
		maxV = max(maxV, g[i][i+n-1])
		minV = min(minV, f[i][i+n-1])
	}
	fmt.Println(minV)
	fmt.Println(maxV)
}
