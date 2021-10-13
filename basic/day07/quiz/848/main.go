package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	n     int
	h     [N]int
	e, ne [2*N]int
	idx   int
	ans   = N
	st    [N]bool
)

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 返回以u为根节点的树的节点个数
func dfs(u int) int {
	st[u] = true
	var nodeCnt int
	var m int // 去掉该节点的最大连通分量
	nodeCnt++
	for i := h[u]; i != -1; i = ne[i] {
		nodeId := e[i]
		if !st[nodeId] {
			subCnt := dfs(nodeId)
			nodeCnt += subCnt
			m = max(m, subCnt)
		}
	}
	m = max(m, n-nodeCnt)
	ans = min(ans, m)
	return nodeCnt
}

func min(a int, m int) int {
	if a < m {
		return a
	}
	return m
}

func init() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 3*N)
	fmt.Fscanf(reader, "%d\n", &n)
	var a, b int
	for i := 1; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		add(a, b)
		add(b, a)

	}
	dfs(1)
	fmt.Println(ans)
}
