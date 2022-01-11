// https://www.acwing.com/problem/content/801/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	n, res int
	a      [N]int
	m      = make(map[int]int)
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	j := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		m[a[i]]++
		for m[a[i]] > 1 {
			m[a[j]]--
			j++
		}
		res = max(res, i-j+1)
	}
	fmt.Println(res)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
