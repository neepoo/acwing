package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	f    [N]int
	t, m int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &t, &m)
	var v, w int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &v, &w)
		for j := t; j >= v; j-- {
			f[j] = max(f[j], f[j-v]+w)
		}
	}
	fmt.Println(f[t])
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
