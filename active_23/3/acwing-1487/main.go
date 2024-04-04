package main

import (
	"bufio"
	. "fmt"
	"os"
)

const (
	N   = 1e5 + 10
	MOD = 1e9 + 7
)

var (
	n1, n2, m int
	a, b      [N]int
	f         [N]int
)

func main() {
	f[0] = 1 // 什么都不选就是一种方案
	r := bufio.NewReader(os.Stdin)
	Fscan(r, &n1, &n2, &m)
	for i := 0; i < n1; i++ {
		Fscan(r, &a[i+1])
	}
	for i := 0; i < n2; i++ {
		Fscan(r, &b[i+1])
	}
	// 完全背包
	for i := 1; i <= n1; i++ {
		for j := a[i]; j <= m; j++ {
			f[j] = (f[j] + f[j-a[i]]) % MOD
		}
	}
	// 01 背包
	for i := 1; i <= n2; i++ {
		for j := m; j >= b[i]; j-- {
			f[j] = (f[j] + f[j-b[i]]) % MOD
		}
	}
	Println(f[m])
}
