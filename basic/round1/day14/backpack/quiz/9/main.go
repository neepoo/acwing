package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 110

var (
	v, w [N][N]int
	f, s [N]int
	n, m int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 40000)
	Fscanf(reader, "%d %d\n", &n, &m)
	for i := 1; i <= n; i++ {
		var si int
		Fscanf(reader, "%d\n", &si)
		s[i] = si
		var a, b int
		for j := 1; j <= si; j++ {
			Fscanf(reader, "%d %d\n", &a, &b)
			v[i][j] = a
			w[i][j] = b
		}
	}
	for i := 1; i <= n; i++ {
		for j := m; j >= 0; j-- {
			for k := 1; k <= s[i]; k++ {
				if j >= v[i][k] {
					f[j] = max(f[j], f[j-v[i][k]]+w[i][k])
				}
			}
		}
	}
	Println(f[m])
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
