// https://www.acwing.com/problem/content/832/

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var n int

func main() {
	l := list.New()
	var x int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		for l.Len() > 0 && l.Back().Value.(int) >= x {
			l.Remove(l.Back())
		}
		if l.Len() == 0 {
			fmt.Printf("%d ", -1)
		} else {
			fmt.Printf("%d ", l.Back().Value.(int))
		}
		l.PushBack(x)
	}
}
