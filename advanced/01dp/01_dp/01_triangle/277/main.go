package main
// https://www.acwing.com/solution/content/12389/  证明
// https://www.acwing.com/problem/content/277/  题目
import (
	"bufio"
	"fmt"
	"os"
)

const N = 55

var (
	g [N][N]int
	f [2 * N][N][N]int

	m, n int
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

func dp() {
	for k := 2; k <= m+n; k++ {
		for x1 := max(k-n, 1); x1 <= min(m, k-1); x1++ {
			for x2 := max(k-n, 1); x2 <= min(m, k-1); x2++ {
				y1, y2 := k-x1, k-x2
				t := g[x1][y1]
				if x1 != x2 {
					// 两条路径走到的不是同一个点
					t += g[x2][y2]
				}
				f[k][x1][x2] = max(f[k][x1][x2], f[k-1][x1-1][x2-1]+t)
				f[k][x1][x2] = max(f[k][x1][x2], f[k-1][x1][x2-1]+t)
				f[k][x1][x2] = max(f[k][x1][x2], f[k-1][x1-1][x2]+t)
				f[k][x1][x2] = max(f[k][x1][x2], f[k-1][x1][x2]+t)
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &m, &n)
	for j := 1; j <= m; j++ {
		for i := 1; i <= n; i++ {
			if i == n {
				fmt.Fscanf(reader, "%d \n", &g[j][i])
			} else {
				fmt.Fscanf(reader, "%d ", &g[j][i])
			}
		}
	}
	dp()
	fmt.Println(f[m+n][m][m])
}
