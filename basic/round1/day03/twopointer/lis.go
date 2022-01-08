// 最长上升子序列,枚举右端点
// https://www.acwing.com/problem/content/801/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 100010

var (
	n int
	a [N]int
	s [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lis(q []int) int {
	var res int
	j := 0
	for i := 0; i < n; i++ {
		s[q[i]]++
		for j < n && s[q[i]] > 1 {
			s[q[j]]--
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	bufio.NewReaderSize(reader, 4*2*N)
	Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		Fscanf(reader, "%d ", &a[i])
	}
	Println(lis(a[:n]))
}
