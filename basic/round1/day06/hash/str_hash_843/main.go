package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N        = 100010
	P uint64 = 131
)

var (
	n, m int
	s    string
	p    [N]uint64
	h    [N]uint64
)

func get(l, r int) uint64 {
	return h[r] - h[l-1]*p[r-l+1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 4*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	fmt.Fscanf(reader, "%s\n", &s)
	p[0] = 1
	for idx, ch := range s {
		p[idx+1] = p[idx] * P
		h[idx+1] = h[idx]*P + uint64(ch-'a'+1)
	}
	var l1, r1, l2, r2 int
	for m > 0 {
		fmt.Fscanf(reader, "%d %d %d %d\n", &l1, &r1, &l2, &r2)
		if get(l1, r1) == get(l2, r2) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		m--
	}
}
