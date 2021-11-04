package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = 0x3f3f3f3f
	N   = 1500
	M   = 3000
)

var (
	n, p, m int

	id      [N]int

	h        [N]int
	ne, e, w [M]int
	idx  int

	dist [N]int
	q    [N]int

	st   [N]bool
)

func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func init() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
}

func spfa(start int) int {
	// 初始化距离
	for i := 0; i < N; i++ {
		dist[i] = INF
	}
	dist[start] = 0
	hh, tt := 0, 1
	q[0] = start
	st[start] = true
	for hh != tt {
		t := q[hh]
		st[t] = false
		hh++
		if hh == N {
			hh = 0
		}
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[t]+w[i] {
				dist[j] = dist[t] + w[i]
				if !st[j] {
					q[tt] = j
					tt++
					if tt == N {
						tt = 0
					}
					st[j] = true
				}
			}
		}

	}
	var res = 0
	for i := 0; i < n; i++ {
		j := id[i]
		if dist[j] == INF {
			return INF
		}
		res += dist[j]
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2 * N + 1500 * 6)
	fmt.Fscanf(reader, "%d %d %d\n", &n, &p, &m)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &id[i])
	}
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		add(a, b, c)
		add(b, a, c)
	}
	var res = INF
	for i := 1; i <= p; i++ {
		res = min(res, spfa(i))
	}
	fmt.Println(res)
}

func min(res int, i int) int {
	if res < i {
		return res
	}
	return i
}
