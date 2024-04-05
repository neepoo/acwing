package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	N   = 15e4 + 10
	MAX = 0x3f3f3f3f
)

var (
	n, m        int
	h, e, ne, w [N]int
	idx         int
	st          [N]bool
	dist        [N]int
)

func add(a, b, c int) {
	e[idx] = b
	ne[idx] = h[a]
	w[idx] = c
	h[a] = idx
	idx++
}

type node struct {
	d, v int
}

type pq []node

func (p *pq) Len() int {
	return len(*p)
}

func (p *pq) Less(i, j int) bool {
	return (*p)[i].d < (*p)[j].d
}

func (p *pq) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *pq) Push(x interface{}) {
	*p = append(*p, x.(node))
}

func (p *pq) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func dijkstra() {
	dist[1] = 0
	q := make(pq, 0)
	heap.Init(&q)
	heap.Push(&q, node{0, 1})
	for q.Len() > 0 {
		head := heap.Pop(&q).(node)
		if st[head.v] {
			continue
		}
		st[head.v] = true
		for j := h[head.v]; j != -1; j = ne[j] {
			v1, w1 := e[j], w[j]
			if dist[head.v]+w1 < dist[v1] {
				dist[v1] = dist[head.v] + w1
				heap.Push(&q, node{dist[v1], v1})
			}
		}
	}
	d := dist[n]
	if d == MAX {
		fmt.Println(-1)
	} else {
		fmt.Println(d)
	}
}

func init() {
	for i := 0; i < N; i++ {
		dist[i] = MAX
		h[i] = -1
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
	dijkstra()
}
