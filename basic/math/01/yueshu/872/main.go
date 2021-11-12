/*
https://www.acwing.com/problem/content/872
约数个数
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = int(1e9 + 7)

var (
	n int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var x int
	var ans = 1
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &x)
		for i := 2; i <= x/i; i++ {
			for x%i == 0 {
				m[i]++
				x /= i
			}
		}
		if x > 1 {
			m[x]++
		}
	}
	for _, cnt := range m {
		ans = ans * (cnt + 1) % mod
	}
	fmt.Println(ans)
}
