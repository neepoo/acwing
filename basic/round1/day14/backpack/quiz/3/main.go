// https://www.acwing.com/problem/content/3/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1010

var (
	f    [N]int
	v, w [N]int
	n, m int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 4*N)
	Fscanf(reader, "%d %d\n", &n, &m)
	var a, b int
	for i := 1; i <= n; i++ {
		Fscanf(reader, "%d %d\n", &a, &b)
		v[i] = a
		w[i] = b
	}

	for i := 1; i <= n; i++ {
		for j := v[i]; j <= m; j++ {
			f[j] = max(f[j], f[j-v[i]]+w[i])
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
