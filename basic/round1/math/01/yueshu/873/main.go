/*
https://www.acwing.com/problem/content/873
约数之和
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
	for p, cnt := range m {
		var tmp = 1
		for cnt > 0 {
			tmp = (tmp*p + 1) % mod
			cnt--
		}
		ans = ans * tmp % mod
	}
	fmt.Println(ans)
}
