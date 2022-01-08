package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N    = 200003
	null = 0x3f3f3f3f
)

var (
	n int
	q [N]int
)

func init() {
	for i := 0; i < N; i++ {
		q[i] = null
	}
}

func find(x int) int {
	k := (x%N + N) % N
	for q[k] != x && q[k] != null { // 坑位有人,并且这个人不是自己那么就继续往下找
		k++
		if k == N {
			k = 0
		}
	}
	return k
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	fmt.Fscanf(reader, "%d\n", &n)
	var op string
	var x int
	for n > 0 {
		fmt.Fscanf(reader, "%s %d\n", &op, &x)
		k := find(x)
		switch op {
		case "I":
			q[k] = x
		case "Q":
			if q[k] == null {
				fmt.Println("No")
			} else {
				fmt.Println("Yes")
			}
		}
		n--
	}
}
