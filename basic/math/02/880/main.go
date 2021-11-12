package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n       int
	a, b, m int
)

func gcm(a, b int, x, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	var d, x1, y1 int
	d = gcm(b, a%b, &x1, &y1)
	*x = y1
	*y = x1 - (a/b)*y1
	return d
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &m)
		var x, y int
		d := gcm(a, m, &x, &y)
		if b%d != 0 {
			fmt.Println("impossible")
		} else {
			fmt.Println(x * b / d % m)
		}
	}
}
