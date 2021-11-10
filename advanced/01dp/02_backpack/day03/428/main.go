/*
https://www.acwing.com/problem/content/description/428/
01 luoti
 */
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 30010
)

var (
	n, m int
	v, p [N]int
	f    [N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 1; i <= m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &v[i], &p[i])
	}
	for i:=1;i<=m;i++{
		for j:=n;j>=v[i];j--{
			f[j] = max(f[j], f[j-v[i]] + v[i]*p[i])
		}
	}
	fmt.Println(f[n])
}


func max(a, b int) int {
	if a > b{
		return a
	}
	return b
}