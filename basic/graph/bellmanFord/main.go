package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	a, b, c int
}

const (
	N   = 510
	MAX = 0x3f3f3f3f
)

var (
	n, m, k      int
	dist, backup [N]int
	edges        []edge
)

func init() {
	for i := 0; i < N; i++ {
		dist[i] = MAX
	}
}

func bellmanFord() {
	dist[1] = 0
	for i := 0; i < k; i++ {
		// 对所有的边进行relax
		copy(backup[:], dist[:])
		for j := range edges {
			e := edges[j]
			if dist[e.b] > backup[e.a]+e.c {
				dist[e.b] = backup[e.a] + e.c
			}
		}
	}
	d := dist[n]
	if d >= MAX/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(d)
	}
}

// https://www.acwing.com/problem/content/855/
func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m, &k)
	edges = make([]edge, m)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		edges[i] = edge{a, b, c}
	}
	bellmanFord()
}
