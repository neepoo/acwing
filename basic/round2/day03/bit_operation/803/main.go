//https://www.acwing.com/problem/content/803/
package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func cnt(x int) int {
	res := 0
	for x > 0 {
		x -= lowbit(x)
		res++
	}
	return res
}

func lowbit(x int) int {
	return x & (-x)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		fmt.Print(cnt(x), " ")
	}
}
