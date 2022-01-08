package main

import (
	"container/list"
	"fmt"
)

func lastRemaining(n int, m int) int {
	l := list.New()
	for i := 0; i < n; i++ {
		l.PushBack(i)
	}
	cnt := 1
	e := l.Front()
	for l.Len() > 1 {
		ne := e.Next()
		if cnt%m == 0 {
			l.Remove(e)
		}
		e = ne

		if e == nil {
			e = l.Front()
		}
		cnt++
	}
	return l.Front().Value.(int)
}

func main() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	//l.Remove(l.Front())
	//l.Remove(l.Front())
	// Iterate through list and print its contents.

	for e := l.Front(); ; e = e.Next() {
		if e == nil {
			e = l.Front()
		}

		fmt.Println(e.Value)
	}

	// 2 0 4 1
	fmt.Println(lastRemaining(5, 3))
	// 70866
	//116922
	//fmt.Println(lastRemaining(116922, 70866))
}
