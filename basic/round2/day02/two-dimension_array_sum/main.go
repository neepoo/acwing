// https://www.acwing.com/problem/content/798/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	n, m, q        int
	x1, y1, x2, y2 int
	a              [N][N]int
	b              [N][N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m, &q)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			b[i][j] = a[i][j] + b[i][j-1] + b[i-1][j] - b[i-1][j-1]
		}
	}
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		fmt.Println(query(x1, y1, x2, y2))
	}
}

func query(x1 int, y1 int, x2 int, y2 int) int {
	res := b[x2][y2] - b[x2][y1-1] - b[x1-1][y2] + b[x1-1][y1-1]
	return res
}
