// https://www.acwing.com/problem/content/831/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	q  [N]int
	hh = 0
	tt = -1
)

func push(x int) {
	tt++
	q[tt] = x
}

func pop() {
	hh++
}

func empty() string {
	if hh <= tt {
		return "NO"
	}

	return "YES"
}

func query() int {
	return q[hh]
}

func main() {
	var (
		n, x int
		op   string
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &op)
		switch op {
		case "push":
			fmt.Fscan(reader, &x)
			push(x)
		case "pop":
			pop()
		case "empty":
			fmt.Println(empty())
		case "query":
			fmt.Println(query())
		}
	}
}
