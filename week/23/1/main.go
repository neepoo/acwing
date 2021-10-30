package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	ans = int(-1e8)
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &x)
		if x < 0 {
			ans = max(ans, x)
			continue
		}

		flag := false
		for j := 0; j*j <= x; j++ {
			if j*j == x {
				flag = true
				break
			}
		}
		if !flag {
			ans = max(ans, x)
		}
	}
	fmt.Println(ans)
}

func max(a int, x int) int {
	if a > x {
		return a
	}
	return x
}
