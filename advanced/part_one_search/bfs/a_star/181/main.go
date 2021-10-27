package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

const end = "12345678x"

type father struct {
	val string // 父亲状态的字符串表示
	op  string // 父亲状态经过 op 操作到达此状态
}

func newFather(val string, op string) father {
	return father{val: val, op: op}
}

type item struct {
	val string // 当前状态的字符串表示
	sd  int    // 源点到当前状态的距离
	dd  int    // 当前状态到目标点的估计距离
}

func newItem(val string, sd int, dd int) item {
	return item{val: val, sd: sd, dd: dd}

}

type Pq []item

func (p Pq) Len() int {
	return len(p)
}

func (p Pq) Less(i, j int) bool {
	return p[i].sd+p[i].dd < p[j].sd+p[j].dd
}

func (p Pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pq) Push(x interface{}) {
	it := x.(item)
	*p = append(*p, it)
}

func (p *Pq) Pop() interface{} {
	old := *p
	n := len(old)
	it := old[n-1]
	*p = old[:n-1]
	return it
}

var (
	start string
	prev  = make(map[string]father)
	// 四个方向 老样子
	dx   = []int{-1, 0, 1, 0}
	dy   = []int{0, 1, 0, -1}
	ops  = []string{"u", "r", "d", "l"}
	dist = make(map[string]int) // 记录当前状态到源点的距离
)

// f 返回s状态到目标状态的估计值
func f(s string) int {
	res := 0
	for i := 0; i < 9; i++ {
		if s[i] != 'x' {
			t := s[i] - '1'
			// 曼哈顿距离
			res += abs(i/3-int(t/3)) + abs(i%3-int(t%3))
		}
	}
	return res
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -1 * i
}

func bfs() {
	q := make(Pq, 0)
	heap.Init(&q)
	// 源点入队
	heap.Push(&q, newItem(start, 0, f(start)))
	dist[start] = 0

	for q.Len() > 0 {
		// 获取优先队列头
		headInterface := heap.Pop(&q)
		head := headInterface.(item)
		// 获取当前状态的字符串表示
		currentState := head.val
		// 获取当前状态到源点的距离
		step := dist[currentState]
		// a star 终止条件
		if currentState == end {
			break
		}

		x, y := -1, -1
		// 找到'x'在当前状态中的坐标
		for i := 0; i < 9; i++ {
			if currentState[i] == 'x' {
				x = i / 3
				y = i % 3
				break
			}
		}
		state := currentState
		currentR := []rune(state)
		// 枚举四个方向
		for i := 0; i < 4; i++ {
			a := x + dx[i]
			b := y + dy[i]
			// 需要在合法的坐标范围内
			if a >= 0 && a < 3 && b >=0 && b < 3{
				currentR[a * 3 + b], currentR[x * 3 + y] = currentR[x * 3 + y], currentR[a * 3 + b]
				d, ok := dist[string(currentR)]
				if !ok || step + 1 < d{
					dist[string(currentR)] = step+1
					prev[string(currentR)] = newFather(currentState, ops[i])
					heap.Push(&q, newItem(string(currentR), dist[string(currentR)], f(string(currentR))))
				}
				// 恢复现场
				currentR[a * 3 + b], currentR[x * 3 + y] = currentR[x * 3 + y], currentR[a * 3 + b]
			}
		}
	}
	var res string
	e := end
	for e != start{
		res += prev[e].op
		e = prev[e].val
	}
	for i := len(res) - 1; i >=0; i-- {
		fmt.Print(string(res[i]))
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, " ")
	start = strings.Join(tokens, "")
	nxd := 0
	// 逆序对数量如果是偶数,才可能有解
	for i := 0; i < 9; i++ {
		if start[i] == 'x' {
			continue
		}
		for j := i + 1; j < 9; j++ {
			if start[i] > start[j] {
				nxd++
			}
		}
	}
	if nxd&1 == 1 {
		fmt.Println("unsolvable")
		return
	}
	bfs()
}
