package main

import (
	"bufio"
	"fmt"
	"os"
)

const V = 20010

var (
	f    [V]int
	v, n int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &v)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscanf(reader, "%d\n", &a)

		for j := v; j >= a; j-- {
			f[j] = max(f[j], f[j-a] + a)
		}
	}
	fmt.Println(v - f[v])
}

func max(i int, i2 int) int {
	if i > i2{
		return i
	}
	return i2
}
