package main

import (
	"bufio"
	"fmt"
	"os"
)

//https://www.acwing.com/problem/content/798/
const N = 1010

var (
	n, m, q int
	a       [N][N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*1024*1024)
	fmt.Fscanf(reader, "%d %d %d\n", &n, &m, &q)
	//fmt.Println(n, m, q)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscanf(reader, "%d ", &a[i][j])
			// 计算前缀和矩阵
			a[i][j] += a[i-1][j] + a[i][j-1] - a[i-1][j-1]
		}
	}
	var x1, y1, x2, y2 int
	for i := 0; i < q; i++ {
		fmt.Fscanf(reader, "%d %d %d %d\n", &x1, &y1, &x2, &y2)
		fmt.Println(a[x2][y2] - a[x2][y1-1] - a[x1-1][y2] + a[x1-1][y1-1])
	}
}
