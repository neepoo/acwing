package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	n, m, q           int
	x1, y1, x2, y2, c int
	a                 [N][N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m, &q)
	var x int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(reader, &x)
			add(i, j, i, j, x)
		}
	}
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &x1, &y1, &x2, &y2, &c)
		add(x1, y1, x2, y2, c)
	}
	// sum
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			a[i][j] += a[i][j-1] + a[i-1][j] - a[i-1][j-1]
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}

func add(x1 int, y1 int, x2 int, y2 int, c int) {
	a[x1][y1] += c
	a[x2+1][y1] -= c
	a[x1][y2+1] -= c
	a[x2+1][y2+1] += c
}
