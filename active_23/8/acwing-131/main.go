package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 10

var (
	h    [N]int
	n    int
	l, r [N]int
	stk  [N]int
	tt   int
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func run() {
	h[0] = -1
	h[n+1] = -1

	tt = 0
	stk[tt] = 0
	var i int
	for i = 1; i <= n; i++ {
		for tt >= 0 && h[stk[tt]] >= h[i] {
			tt--
		}
		l[i] = stk[tt]
		tt++
		stk[tt] = i
	}

	tt = 0
	stk[tt] = n + 1
	for i = n; i >= 1; i-- {
		for tt >= 0 && h[stk[tt]] >= h[i] {
			tt--
		}
		r[i] = stk[tt]
		tt++
		stk[tt] = i
	}
	var res int64
	for i = 1; i <= n; i++ {
		res = max(res, int64(h[i])*int64(r[i]-l[i]-1))
	}
	fmt.Println(res)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fscan(reader, &n)
		if n == 0 {
			break
		}
		for i := 1; i <= n; i++ {
			fmt.Fscan(reader, &h[i])
		}
		run()
	}

}
