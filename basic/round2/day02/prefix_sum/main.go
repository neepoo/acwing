package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 10

var (
	n, m int
	a    [N]int
	l, r int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
		a[i] += a[i-1]
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &l, &r)
		fmt.Println(a[r] - a[l-1])
	}
}
