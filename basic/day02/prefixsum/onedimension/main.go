package main

import (
	"fmt"
	"os"
)

// https://www.acwing.com/problem/content/797/

const N = 100010

var (
	n, m int
	q    [N]int
	l, r int
)

func main() {
	fmt.Fscanln(os.Stdin, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &q[i])
		q[i] += q[i-1]
	}
	for i:=0; i< m;i++{
		fmt.Scanf("%d %d", &l, &r)
		fmt.Println(q[r] - q[l-1])
	}

}
