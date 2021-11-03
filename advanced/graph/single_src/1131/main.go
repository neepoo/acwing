package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type node struct {
	nodeId int // 节点的id
	d      int // 距离源点的距离
}

type pq []node

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i].d < p[j].d
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x interface{}) {
	nd := x.(node)
	*p = append(*p, nd)
}

func (p *pq) Pop() interface{} {
	old := *p
	n := len(old)
	last := old[n-1]
	*p = old[:n-1]
	return last
}

const (
	N   = 2600     // 点的上限
	M   = 6300 * 2 // 边的上限
	INF = 0x3f3f3f3f
)

var (
	n        int // 点的数量
	m        int // 边的数量
	S        int // 起点
	T        int // 终点
	idx      int
	h        [N]int
	ne, w, e [M]int

	dist [N]int
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
		dist[i] = INF
	}
}

func dijkstra() {
	dist[S] = 0
	for i := 0; i < n; i++ {
		// 找到剩下的节点距离源点最短的
		tt := -1
		for j := 1; j <= n; j++ {
			if !st[j] && (tt == -1 || dist[tt] > dist[j]) {
				tt = j
			}
		}
		st[tt] = true
		for i := h[tt]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[tt]+w[i] {
				dist[j] = dist[tt] + w[i]
			}
		}
	}
}

func dijkstraHeap() {
	p := make(pq, 0)
	heap.Init(&p)
	heap.Push(&p, node{S, 0})
	dist[S] = 0
	for p.Len() > 0 {
		head := heap.Pop(&p)
		hn := head.(node)
		if st[hn.nodeId] {
			continue
		}
		st[hn.nodeId] = true
		// 松弛
		for i := h[hn.nodeId]; i != -1; i = ne[i] {
			nodeID := e[i]
			ds := w[i]

			if dist[nodeID] > dist[hn.nodeId]+ds {
				dist[nodeID] = dist[hn.nodeId] + ds
				heap.Push(&p, node{nodeID, dist[nodeID]})
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, M*3)
	fmt.Fscanf(reader, "%d %d %d %d\n", &n, &m, &S, &T)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		add(a, b, c)
		add(b, a, c)
	}
	dijkstraHeap()
	fmt.Println(dist[T])
}
