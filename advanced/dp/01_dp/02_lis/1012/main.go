package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1010

var (
	n       int
	a, f, g [N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, " ")
	n = len(tokens)
	var ans int
	for idx, s := range tokens {
		v, _ := strconv.Atoi(s)
		a[idx+1] = v
	}
	for i := 1; i <= n; i++ {
		f[i] = 1
		for j := 1; j < i; j++ {
			if a[j] >= a[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	fmt.Println(ans)

	var cnt int // 序列的总数
	// g 是单调上升的
	for i := 1; i <= n; i++ {
		var k int // 表示第几个序列
		for k < cnt && g[k] < a[i] {
			k++
		}
		g[k] = a[i]
		if k >= cnt {
			cnt++
		}

	}
	fmt.Println(cnt)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
