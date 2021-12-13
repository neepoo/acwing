package main

import (
	"bufio"
	. "fmt"
	"os"
)

type query struct {
	x  int
	y  int
	op int
}
type querys []query

const N = 200010

var (
	t, n, cnt int
	q         [N]int
	m         map[int]int
)

// 离散化
func get(x int) int {
	if m[x] == 0 {
		cnt++
		m[x] = cnt
	}
	return m[x]
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 1<<24)
	Fscan(reader, &t)
	for t > 0 {
		Fscan(reader, &n)
		for i := 0; i < N; i++ {
			q[i] = i
		}
		var x, y, op int
		qs := make(querys, n)
		m = make(map[int]int)
		cnt = 0
		for i := 0; i < n; i++ {
			Fscan(reader, &x, &y, &op)
			x = get(x)
			y = get(y)
			qs[i] = query{
				x:  x,
				y:  y,
				op: op,
			}
		}
		// merge
		for i := 0; i < n; i++ {
			if qs[i].op == 1 {
				pa, pb := find(qs[i].x), find(qs[i].y)
				q[pa] = q[pb]
			}
		}
		var hasConflict bool
		for i := 0; i < n; i++ {
			if qs[i].op == 0 {
				pa, pb := find(qs[i].x), find(qs[i].y)
				if pa == pb {
					hasConflict = true
					break
				}
			}
		}
		if hasConflict {
			Println("NO")
		} else {
			Println("YES")
		}
		t--
	}

}

func find(x int) int {
	if x != q[x] {
		q[x] = find(q[x])
	}
	return q[x]
}
