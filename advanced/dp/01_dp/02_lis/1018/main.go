package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	ans  int
	a, f [N]int
	n    int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &a[i])
		ans = max(ans, a[i])
	}
	for i := 1; i <= n; i++ {
		f[i] = a[i]
		for j := 1; j < i; j++ {
			if a[i] > a[j] {
				f[i] = max(f[i], f[j]+a[i])
				ans = max(ans, f[i])
			}
		}
	}
	fmt.Println(ans)
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}
