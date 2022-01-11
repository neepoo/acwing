// https://www.acwing.com/problem/content/830/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	stk [N]int
	tt  int
)

func push(x int) {
	tt++
	stk[tt] = x
}

func pop() int {
	res := stk[tt]
	tt--
	return res
}

func empty() bool {
	return tt == 0
}
func query() int {
	return stk[tt]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	var op string
	var x int
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
			emp := empty()
			if emp {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case "query":
			fmt.Println(query())
		}
	}
}
