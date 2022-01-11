// https://www.acwing.com/problem/content/828/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	n     int
	idx   int
	e, ne [N]int
	head  int
	op    string
)

func init() {
	head = -1
	idx = 0
	for i := 0; i < N; i++ {
		ne[i] = -1
	}
}

func insertHead(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

func deleteHead() {
	head = ne[head]
}

func insertBack(k, x int) {
	e[idx] = x
	ne[idx] = ne[k]
	ne[k] = idx
	idx++
}

func remove(k int) {
	if k == -1 {
		deleteHead()
	} else {
		ne[k] = ne[ne[k]]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	var k, x int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &op)
		switch op {
		case "H":
			fmt.Fscan(reader, &x)
			insertHead(x)
		case "D":
			fmt.Fscan(reader, &k)
			remove(k - 1)
		case "I":
			fmt.Fscan(reader, &k, &x)
			insertBack(k-1, x)
		}
	}
	for i := head; i != -1; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
}
