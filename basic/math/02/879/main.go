package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, a, b int
)

func exGcm(a, b int, x, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	var x1, y1, d int
	// 先求出下一层的系数
	d = exGcm(b, a%b, &x1, &y1)
	// 用来更新当前层的系数
	*x = y1
	*y = x1 - a/b*y1
	return d
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var x, y int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		exGcm(a, b, &x, &y)
		fmt.Println(x, y)
	}
}
