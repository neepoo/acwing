// https://www.acwing.com/problem/content/849/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010
const MIN = -1

var (
	n, m int

	idx      int
	e, ne, h [N]int

	d  [N]int
	q  [N]int
	hh = 0
	tt = -1
)

func init() {
	for i := 0; i < N; i++ {
		d[i] = MIN
		h[i] = -1
	}
}

func bfs() {
	tt++
	q[tt] = 1
	d[1] = 0
	for hh <= tt {
		head := q[hh]
		hh++
		for i := h[head]; i != -1; i = ne[i] {
			nodeID := e[i]
			// 这里不需要st数组的原因是. d[i]=MAX暗含了该节点并未访问过
			if d[nodeID] == MIN {

				d[nodeID] = d[head] + 1
				tt++
				q[tt] = nodeID
			}
		}
	}
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 3*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		add(a, b)
	}
	bfs()
	fmt.Println(d[n])
}
