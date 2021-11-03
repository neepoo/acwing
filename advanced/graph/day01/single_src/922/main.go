// https://www.acwing.com/problem/content/922/
// 先计算坐车数量 然后减1得到换乘数量

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	N = 510    // 公交站上限
	M = 250000 // 边上线
)

var (
	m, n int

	h     [N]int
	ne, e [M]int
	idx   int
	st    [N]bool
)

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

type pair struct {
	id int
	ds int
}

func bfs() int {
	var q [N]pair
	hh, tt := 0, 1
	q[0] = pair{1, 0}
	st[1] = true
	if 1 == n{
		return 0
	}
	for hh != tt {
		head := q[hh]
		//fmt.Printf("head=%#v\n", head)
		hh++
		if hh == N{
			hh = 0
		}
		for i := h[head.id]; i != -1; i = ne[i] {
			j := e[i]
			//fmt.Println("j=", j, "n=", n)
			if j == n {
				return head.ds + 1
			}
			if !st[j] {
				q[tt] = pair{j, head.ds + 1}
				tt++
				if tt == N{
					tt = 0
				}
				st[j] = true
			}
		}
	}
	return -1
}

func main() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
//	reader := strings.NewReader(`3 7
//6 7
//4 7 3 6
//2 1 3 5`)
	reader:=bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	l := scanner.Text()
	ts := strings.Split(l, " ")
	m, _ = strconv.Atoi(ts[0])
	n, _ = strconv.Atoi(ts[1])
	for i:=0;i<m;i++ {
		scanner.Scan()
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		tmp := make([]int, len(tokens))
		for inx, s := range tokens {
			i, _ := strconv.Atoi(s)
			tmp[inx] = i
		}
		for i := 0; i < len(tmp); i++ {
			for j := i + 1; j < len(tmp); j++ {
				add(tmp[i], tmp[j])
			}
		}
	}
	res := bfs()

	switch {
	case res == -1:
		fmt.Println("NO")
	case res == 0:
		fmt.Println(0)
	default:
		fmt.Println(res - 1)
	}
}
