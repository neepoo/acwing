// https://www.acwing.com/problem/content/829/
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var (
	n, k, x int
	op      string
	idx     = 1

	m = make(map[int]*list.Element)
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	l := list.New()
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &op)
		switch op {
		case "L":
			fmt.Fscan(reader, &x)
			ele := l.PushFront(x)
			m[idx] = ele
			idx++
		case "R":
			fmt.Fscan(reader, &x)
			ele := l.PushBack(x)
			m[idx] = ele
			idx++
		case "D":
			fmt.Fscan(reader, &k)
			ele := m[k]
			l.Remove(ele)
		case "IL":
			fmt.Fscan(reader, &k, &x)
			ele := m[k]
			newEle := l.InsertBefore(x, ele)
			m[idx] = newEle
			idx++
		case "IR":
			fmt.Fscan(reader, &k, &x)
			ele := m[k]
			newEle := l.InsertAfter(x, ele)
			m[idx] = newEle
			idx++
		}
	}
	head := l.Front()
	for head != nil {
		fmt.Printf("%d ", head.Value.(int))
		head = head.Next()
	}
}
