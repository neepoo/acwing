package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	n, m, size int
	h          [N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	fmt.Fscanf(reader, "%d %d \n", &n, &m)
	size = n
	x := 0
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d ", &x)
		h[i] = x
	}
	// 初始化堆
	for i := n / 2; i >= 1; i-- {
		down(i)
	}
	for m > 0 {
		fmt.Printf("%d ", h[1])
		h[1] = h[size]
		size--
		down(1)
		m--
	}
}

func down(u int) {
	t := u // 三个节点中最小值的下标
	if 2*u <= size && h[2*u] < h[t] {
		t = 2 * u
	}
	if 2*u+1 <= size && h[2*u+1] < h[t] {
		t = 2*u + 1
	}
	if t != u {
		h[u], h[t] = h[t], h[u]
		down(t)
	}
}
