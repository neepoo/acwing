package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	n    int
	a, f [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &a[i])
	}

	var ans int
	for i := 0; i < n; i++ {
		l, r := 0, ans
		for l < r {
			mid := (l + r + 1) / 2
			// 在 g[l:ans+1] 中找到比a[i]小的最大的数
			if f[mid] < a[i] {
				l = mid
			} else {
				r = mid - 1
			}
		}
		// a[i] 一定小于f[r+1] 并且f数组单调增
		f[r+1] = a[i]
		ans = max(ans, r+1)
	}
	fmt.Println(ans)
}
