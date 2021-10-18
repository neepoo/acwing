package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 1010

var (
	n int
	a [N]int
	f [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		if i == n {
			Fscanf(reader, "%d\n", &a[i])
		} else {
			Fscanf(reader, "%d ", &a[i])
		}
	}

	for i := 1; i <= n; i++ {
		f[i] = 1
		for j := 1; j < n; j++ {
			if a[i] > a[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	var res = f[1]
	for i := 2; i <= n; i++ {
		res = max(res, f[i])
	}
	Println(res)
}
