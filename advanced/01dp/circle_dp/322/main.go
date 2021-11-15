package main
/*
https://www.acwing.com/problem/content/322/
 */

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = 0x3f3f3f3f
	N   = 210
)

var (
	n int
	w [N]int
	f [N][N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &w[i])
		w[i+n] = w[i]
	}
	for length := 3; length <= n+1; length++ {
		for l := 1; l+length-1 <= 2*n; l++ {
			r := l + length - 1
			for k := l + 1; k < r; k++ {
				f[l][r] = max(f[l][r], f[l][k]+f[k][r]+w[l]*w[k]*w[r])
			}
		}
	}
	var res = -INF
	for i := 1; i <= n; i++ {
		res = max(res, f[i][i+n])
	}
	fmt.Println(res)
}
