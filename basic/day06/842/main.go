package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	N = 100003
)

var (
	q     [N]int
	e, ne [N]int
	idx   int
	n     int
)

func find(x int) bool {
	k := (x%N + N) % N
	for i := q[k]; i != -1; i = ne[i] {
		if e[i] == x {
			return true
		}
	}
	return false
}

func insert(x int) {
	k := (x%N + N) % N  // hash value, 防止出现负数
	e[idx] = x
	ne[idx] = q[k]
	q[k] = idx
	idx++
}

func init() {
	for i := 0; i < N; i++ {
		q[i] = -1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 3*N)
	fmt.Fscanf(reader, "%d\n", &n)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		op := tokens[0]
		x, _ := strconv.Atoi(tokens[1])
		switch op {
		case "I":
			insert(x)
		case "Q":
			if !find(x) {
				fmt.Println("No")
			} else {
				fmt.Println("Yes")
			}
		}
	}
}
