package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	N = 1010
	M = 200010
)

type item struct {
	sd     int // 源点到当前点距离
	dd     int // 当前点到终点的估计(实际)距离
	nodeID int // 节点的唯一编号
}

func newItem(sd int, dd int, nodeID int) item {
	return item{sd: sd, dd: dd, nodeID: nodeID}
}

type PriorityQueue []item

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].sd+p[i].dd < p[j].sd+p[j].dd
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	it := x.(item)
	*p = append(*p, it)
}

func (p *PriorityQueue) Pop() interface{} {
	n := len(*p)
	old := *p
	*p = old[:n-1]
	return old[n-1]
}

var (
	h, rh    [N]int
	ne, e, w [M]int
	idx      int

	dist [N]int // 当前点到终点的最短距离
	cnt  [N]int
	st   [N]bool // dijkstra时判断该点是否已经拓展过

	n, m int
	T    int
	K    int
	S    int
)

func init() {
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
		h[i] = -1
		rh[i] = -1
	}
}

func add(h []int, a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 600000)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		add(h[:], a, b, c)
		add(rh[:], b, a, c)
	}
	fmt.Fscanf(reader, "%d %d %d\n", &S, &T, &K)
	if S == T {
		K++
	}
	dijkstra()
	fmt.Println(aStar())
}

func aStar() int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	// 源点入队
	heap.Push(&pq, newItem(0, dist[S], S))
	for pq.Len() > 0 {
		head := heap.Pop(&pq).(item)
		cnt[head.nodeID]++
		if cnt[T] == K {
			return head.sd
		}
		for i := h[head.nodeID]; i != -1; i = ne[i] {
			nodeID := e[i]
			if cnt[nodeID] < K {
				heap.Push(&pq, newItem(head.sd+w[i], dist[nodeID], nodeID))
			}
		}
	}
	return -1
}

func dijkstra() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	// 终点入队
	heap.Push(&pq, newItem(0, 0, T))
	dist[T] = 0
	for pq.Len() > 0 {
		head := heap.Pop(&pq).(item)
		// dijkstra 都是出队才去重复
		if st[head.nodeID] {
			continue
		}
		st[head.nodeID] = true
		for i := rh[head.nodeID]; i != -1; i = ne[i] {
			nodeID := e[i]
			// 松弛操作
			if dist[nodeID] > dist[head.nodeID]+w[i] {
				dist[nodeID] = dist[head.nodeID] + w[i]
				// 入队
				heap.Push(&pq, newItem(dist[nodeID], 0, nodeID))
			}
		}
	}
}
