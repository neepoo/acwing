package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 3010

var (
	n int // 面值的总数
	m int // 目标值

	f [N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	f[0] = 1

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscanf(reader, "%d\n", &x)

		for j := x; j <= m; j++ {
			f[j]+=f[j-x]
		}
	}
	fmt.Println(f[m])
}
