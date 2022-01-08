package main

import (
	"bufio"
	. "fmt"
	"os"
)

//https://www.acwing.com/problem/content/803/

const N = 100010

var (
	n int
	a [N]int
)

func lowBit(x int) int {
	return x & (-x)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		Fscanf(reader, "%d ", &a[i])
	}
	for i := 0; i < n; i++ {
		j := a[i]
		res := 0
		for j > 0 {
			j -= lowBit(j)
			res++
		}
		Printf("%d ", res)
	}
}
