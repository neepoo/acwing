//https://www.acwing.com/problem/content/852/
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	N   = 150010
	MAX = 0x3f3f3f3f
)

var (
	h, e, ne, w [N]int
	idx         int
	d           [N]int
	st          [N]bool
	n, m        int
	pq          PriorityQueue
)

// An node is something we manage in a dis queue.
type node struct {
	id  int
	dis int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dis <= pq[j].dis
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func init() {
	for i := 0; i < N; i++ {
		d[i] = MAX
		h[i] = -1
	}
}

func dijkstra() int {
	d[1] = 0
	heap.Push(&pq, &node{id: 1, dis: 0})
	for len(pq) > 0 {
		currentMinNode := heap.Pop(&pq).(*node)
		currentMinNodeID := currentMinNode.id
		if st[currentMinNodeID] {
			continue
		}
		st[currentMinNodeID] = true
		for i := h[currentMinNodeID]; i != -1; i = ne[i] {
			nodeID := e[i]
			if d[nodeID] > d[currentMinNodeID]+w[i] {
				d[nodeID] = d[currentMinNodeID] + w[i]
				heap.Push(&pq, &node{id: nodeID, dis: d[nodeID]})
			}
		}
	}
	if d[n] != MAX {
		return d[n]
	}
	return -1
}

func main() {
	pq = make(PriorityQueue, 0)
	heap.Init(&pq)

	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 4*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var a, b, c int
	for m > 0 {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		add(a, b, c)
		m--
	}
	fmt.Println(dijkstra())
}
