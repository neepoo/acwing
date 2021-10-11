package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acwing.com/problem/content/832/
const N = 100010

var (
	n   int
	stk [N]int
	tt  int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 4*N)
	fmt.Fscanf(reader, "%d\n", &n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &x)
		for tt > 0 && stk[tt] >= x {
			tt--
		}
		if tt == 0 {
			fmt.Printf("%d ", -1)
		} else {
			fmt.Printf("%d ", stk[tt])
		}
		tt++
		stk[tt] = x
	}
}
