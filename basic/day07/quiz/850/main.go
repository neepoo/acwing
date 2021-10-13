//https://www.acwing.com/problem/content/850/
// 有向无环图才有拓扑排序
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	h, e, ne [N]int
	idx      int

	n, m int

	cnt [N]int // 记录每个节点的入度

	q  [N]int // 队列
	hh = 0
	tt = -1
)

func init() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
}
func add(a int, b int) {
	cnt[b]++
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

// 有重边和自环不影响结果
func topSort() {
	for i := 1; i <= n; i++ {
		if cnt[i] == 0 {
			tt++
			q[tt] = i
		}
	}
	for hh <= tt {
		head := q[hh]
		hh++
		for i := h[head]; i != -1; i = ne[i] {
			nodeID := e[i]
			cnt[nodeID]--
			if cnt[nodeID] == 0 {
				tt++
				q[tt] = nodeID
			}
		}
	}
	if tt == n-1 {
		for i := 0; i <= tt; i++ {
			fmt.Printf("%d ", q[i])
		}
	} else {
		fmt.Println(-1)
	}

}
func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 4*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		add(a, b)
	}
	topSort()
}
