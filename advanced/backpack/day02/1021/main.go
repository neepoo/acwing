package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 6010

var (
	n, m int
	f    [510][N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 1; i <= n; i++ { // 枚举物品
		var a, b, c int
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		// enumerate v
		for j := 0; j <= m; j++ {
			// enumerate method
			for k := 0; k <= c; k++ {
				if j >= k*a {
					f[i][j] = max(f[i][j], f[i-1][j-a*k]+b*k)
				}
			}
		}
	}

	fmt.Println(f[n][m])
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
