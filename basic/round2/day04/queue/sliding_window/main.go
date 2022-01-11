// https://www.acwing.com/problem/content/156/

// TODO reimplements it by using container/list
package main

import (
	"bufio"
	//"container/list"
	"fmt"
	"os"
)

const N = 1000010

var (
	n, k, x int
	//l *list.List
	q  [N]int
	a  [N]int
	hh = 0
	tt = -1
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &k)
	for i := 0; i < n; i++ {
		if hh <= tt && q[hh] < i-k+1 {
			hh++
		}
		fmt.Fscan(reader, &x)
		a[i] = x
		for hh <= tt && a[q[tt]] >= x {
			tt--
		}
		tt++
		q[tt] = i
		if i >= k-1 {
			fmt.Printf("%d ", a[q[hh]])
		}
	}
	fmt.Println()
	hh, tt := 0, -1
	for i := 0; i < n; i++ {
		if hh <= tt && q[hh] < i-k+1 {
			hh++
		}
		for hh <= tt && a[q[tt]] <= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		if i >= k-1 {
			fmt.Printf("%d ", a[q[hh]])
		}
	}
}
