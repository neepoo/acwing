//https://www.acwing.com/problem/content/4/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 110

var (
	f       [N][N]int
	v, w, s [N]int
	n, m    int
)

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 4*N)
	Fscanf(reader, "%d %d\n", &n, &m)
	var a, b, c int
	for i := 1; i <= n; i++ {
		Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		v[i] = a
		w[i] = b
		s[i] = c
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= s[i]; k++ {
				if j >= v[i]*k {
					f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
				}
			}
		}
	}
	Println(f[n][m])
}
