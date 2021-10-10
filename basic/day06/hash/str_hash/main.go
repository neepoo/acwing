package main

import (
	"bufio"
	"fmt"
	"os"
)

// 字符串前缀hash
// https://cp-algorithms.com/string/string-hashing.html
/*
 q1: 如何定义: p进制数(p取131/13331), q取pow(2,64)
 q2: 如何计算字串的hash
*/
const (
	N = 100010
	c = 131
)

var (
	n, m int
	h, p [N]uint64
	s    string
)

func get(l, r int) uint64 {
	return h[r] - h[l-1]*p[r-l+1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 8*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	fmt.Fscanf(reader, "%s\n", &s)
	p[0] = 1
	for idx, ch := range s {
		p[idx+1] = p[idx] * c
		h[idx+1] = h[idx]*c + uint64(ch)
	}
	var l1, r1, l2, r2 int
	for m > 0 {
		fmt.Fscanf(reader, "%d %d %d %d\n", &l1, &r1, &l2, &r2)
		m--
		if get(l1, r1) == get(l2, r2) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
