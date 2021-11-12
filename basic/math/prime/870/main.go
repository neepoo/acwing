/*
https://www.acwing.com/problem/content/870/
筛法求素数个数
*/
package main

import "fmt"

const N = 1000010

var (
	n   int
	cnt int
	st  [N]bool
)

func main() {
	fmt.Scan(&n)
	// 埃氏筛法
	for i := 2; i <= n; i++ {
		if !st[i] {
			for j := i; j <= n; j += i {
				st[j] = true
			}
			cnt++
		}
	}
	fmt.Println(cnt)

	// 线性筛法
}
