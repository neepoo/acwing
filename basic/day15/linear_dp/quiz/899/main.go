//https://www.acwing.com/problem/content/899/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1010

var (
	n, m int
	a, b string
	f    [N][N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscanf(reader, "%d %d\n", &n, &m)
	Fscanf(reader, "%s\n%s\n", &a, &b)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1])
			if a[i-1] == b[j-1] {
				f[i][j] = max(f[i][j], f[i-1][j-1]+1)
			}
		}
	}
	Println(f[n][m])
}
