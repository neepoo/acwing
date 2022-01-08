/*
https://www.acwing.com/problem/content/869/
分解质因数
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, x int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &x)
		divide(x)
	}
}

func divide(x int) {
	for i := 2; i <= x/i; i++ {
		if x%i == 0 { // 这里的i一定是一个素数
			var cnt int
			for x%i == 0 {
				x /= i
				cnt++
			}
			fmt.Println(i, cnt)
		}
	}
	if x > 1 {
		fmt.Println(x, 1)
	}
	fmt.Println()
}
