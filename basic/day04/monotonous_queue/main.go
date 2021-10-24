 package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1000010

var (
	n, k int
	a    [N]int
	q    [N]int
)

func main() {
	var x int
	hh := 0
	tt := -1
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 6*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &x)
		a[i] = x
		// 判断窗口是否需要左移
		if tt >= hh && i-k+1 > q[hh] {
			hh++
		}
		for tt >= hh && x <= a[q[tt]] {
			tt--
		}
		tt++
		q[tt] = i
		if i >= k-1 {
			fmt.Printf("%d ", a[q[hh]])
		}
	}
	fmt.Println()
	hh = 0
	tt = -1
	for i := 0; i < n; i++ {
		if tt >= hh && i - k + 1 > q[hh] {hh++}
		for tt >= hh && a[i] >= a[q[tt]]{tt--}
		tt++
		q[tt]=i
		if i >= k - 1{
			fmt.Printf("%d ",a[q[hh]])
		}
	}
}
