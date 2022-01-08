package main

// https://www.acwing.com/problem/content/842/
import (
	"bufio"
	"fmt"
	"os"
)

const N = 100003

var (
	h     [N]int
	e, ne [N]int
	idx   int
	n     int
)

func init() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
}

func find(x int) bool {
	k := (x%N + N) % N
	for i := h[k]; i != -1; i = ne[i] {
		if e[i] == x {
			return true
		}
	}
	return false
}

func insert(x int) {
	k := (x%N + N) % N // hash value, 防止出现负数
	e[idx] = x
	ne[idx] = h[k]
	h[k] = idx
	idx++
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	fmt.Fscanf(reader, "%d\n", &n)
	var op string
	var x int
	for n > 0 {
		fmt.Fscanf(reader, "%s %d\n", &op, &x)
		switch op {
		case "I":
			insert(x)
		case "Q":
			if find(x) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
		n--
	}
}
