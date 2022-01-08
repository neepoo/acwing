//https://www.acwing.com/problem/content/875/
package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &x)
		var res = x
		for j := 2; j <= x/j; j++ {
			if x%j == 0 {
				res = res / j * (j - 1)
			}
			for x%j == 0 {
				x /= j
			}
		}
		if x > 1 {
			res = res / x * (x - 1)
		}
		fmt.Println(res)
	}
}
