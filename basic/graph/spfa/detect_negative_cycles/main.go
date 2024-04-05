package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

const (
	N   = 1000010
	MAX = 0x3f3f3f3f
)

var (
	n, m        int
	dist        [N]int
	cnt         [N]int
	st          [N]bool
	idx         int
	h, e, ne, w [N]int
)

func init() {
	for i := 0; i < N; i++ {
		h[i] = -1
		dist[i] = MAX
	}
}

func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func detectNegativeCycles() {
	q := list.New()
	for i := 1; i <= n; i++ {
		q.PushBack(i)
		st[i] = true
	}
	for q.Len() > 0 {
		head := q.Remove(q.Front()).(int)
		st[head] = false
		for j := h[head]; j != -1; j = ne[j] {
			v1, w1 := e[j], w[j]
			if dist[v1] > dist[head]+w1 {
				dist[v1] = dist[head] + w1
				cnt[v1] = cnt[head] + 1
				if cnt[v1] >= n {
					fmt.Println("Yes")
					return
				}
				if !st[v1] {
					st[v1] = true
					q.PushBack(v1)
				}
			}
		}
	}
	fmt.Println("No")
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b, &c)
		add(a, b, c)
	}
	detectNegativeCycles()
}
