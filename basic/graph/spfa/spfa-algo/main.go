package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 1e5 + 10
	MAX = 0x3f3f3f3f
)

var (
	n, m           int
	h, e, ne, w, d [N]int
	q              [N]int
	idx            int
	st             [N]bool
)

func init() {
	for i := 0; i < N; i++ {
		d[i] = MAX
		h[i] = -1
	}
}

func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func spfa() {
	d[1] = 0
	hh, tt := 0, 0
	tt++
	q[tt] = 1
	st[1] = true
	for tt >= hh {
		head := q[hh]
		hh++
		st[head] = false
		for j := h[head]; j != -1; j = ne[j] {
			v1, w1 := e[j], w[j]
			if d[head]+w1 < d[v1] {
				d[v1] = d[head] + w1
				if !st[v1] {
					st[v1] = true
					tt++
					q[tt] = v1
				}
			}
		}
	}
	dn := d[n]
	if dn >= MAX/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(dn)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		add(a, b, c)
	}
	spfa()
}
