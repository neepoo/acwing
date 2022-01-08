//https://www.acwing.com/problem/content/851/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 510
	MAX = 0x3f3f3f3
)

var (
	g [N][N]int
	d [N]int

	st [N]bool

	n, m int
)

func init() {
	for i := 0; i < N; i++ {
		d[i] = MAX
		for j := 0; j < N; j++ {
			g[i][j] = MAX
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 400000)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for m > 0 {
		var a, b, c int
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		add(a, b, c)
		m--
	}
	dijkstra()
}

func dijkstra() {
	d[1] = 0
	for i := 0; i < n-1; i++ {
		var minNodeIdx = -1
		for j := 1; j <= n; j++ {
			if !st[j] && (minNodeIdx == -1 || d[minNodeIdx] > d[j]) {
				minNodeIdx = j
			}
		}
		st[minNodeIdx] = true
		// 松弛
		for j := 1; j <= n; j++ {
			if d[j] > g[minNodeIdx][j]+d[minNodeIdx] {
				d[j] = g[minNodeIdx][j] + d[minNodeIdx]
			}
		}
	}
	if d[n] == MAX {
		fmt.Println(-1)
	} else {
		fmt.Println(d[n])
	}

}

/*
   d[1] = 0
   for i:=0;i<n-1;i++{
       var minNodeIdx = -1
       for j:=1;j<=n;j++{
           if (!st[j] && (minNodeIdx==-1 || d[minNodeIdx] > d[j])){
               minNodeIdx = j
           }
       }
       st[minNodeIdx]=true
       // 松弛
       for j:=1;j<=n;j++{
           if d[j] > g[minNodeIdx][j] + d[minNodeIdx]{
                d[j] = g[minNodeIdx][j] + d[minNodeIdx]
           }
       }
	}
*/

func add(a int, b int, c int) {
	g[a][b] = min(g[a][b], c)
}

func min(i int, c int) int {
	if i < c {
		return i
	}
	return c
}
