package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 110
const M = 25010

var (
	t int
	a [N]int
	f [M]bool
	n int
)

func clear() {
	for i := 0; i < M; i++ {
		f[i] = false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &t)
	for t > 0 {
		var res int
		fmt.Fscanf(reader, "%d\n", &n)

		for i := 0; i < n; i++ {
			if i == n-1 {
				fmt.Fscanf(reader, "%d \n", &a[i])
			} else {
				fmt.Fscanf(reader, "%d ", &a[i])
			}
		}
		sort.Slice(a[:n], func(i, j int) bool {
			return a[i] < a[j]
		})
		clear()
		f[0] = true
		m := a[n-1]
		for i := 0; i < n; i++ {
			if !f[a[i]] {
				res++
			}
			for j := a[i]; j <= m; j++ {
				f[j] = f[j-a[i]] || f[j]
			}
		}
		fmt.Println(res)
		t--
	}
}
