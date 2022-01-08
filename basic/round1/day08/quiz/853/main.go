//https://www.acwing.com/problem/content/853/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 100010
	MAX = 0x3f3f3f3f
)

var (
	n, m int

	e, ne, w, h [N]int
	idx         int

	d [N]int

	st [N]bool // 当前是否在队列中
)

func init() {
	for i := 0; i < N; i++ {
		h[i] = -1
		d[i] = MAX
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
	q := [N]int{}
	hh, tt := 0, -1

	tt++
	q[tt] = 1
	d[1] = 0
	for hh <= tt {
		head := q[hh]
		st[head] = false
		hh++

		// 松弛临边
		for i := h[head]; i != -1; i = ne[i] {
			nodeID := e[i]

			if d[nodeID] > d[head]+w[i] {
				// 就算不如队列,也要更新距离,这是为什么呢?
				d[nodeID] = d[head] + w[i]
				if !st[nodeID] {
					st[nodeID] = true
					tt++
					q[tt] = nodeID
				}
			}
		}
	}
	if d[n] >= MAX/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(d[n])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 4*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var a, b, c int
	for m > 0 {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		add(a, b, c)
		m--
	}
	spfa()
}
