package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type father struct {
	op   string
	dist int
	src  string
}

func newFather(op string, dist int, src string) father {
	return father{op: op, dist: dist, src: src}
}

const start = "12345678"

var (
	end string
	g   [2][4]int
	m   = make(map[string]father)
	q   [1e6]string
)

// 将一个字符串映射到g数组
func set(s string) {
	for idx, ch := range s {
		var col, row int
		row = idx / 4
		if idx >= 4 {
			col = 3 - (idx % 4)
		} else {
			col = idx % 4
		}
		g[row][col], _ = strconv.Atoi(string(ch))
	}
}

// 将g数组的内容映射为一个字符串
func get() string {
	var res string
	for i := 0; i < 4; i++ {
		res += strconv.Itoa(g[0][i])
	}
	for i := 3; i >= 0; i-- {
		res += strconv.Itoa(g[1][i])
	}
	return res
}

// A 交换g数组的上下两行
func A(s string) string {
	set(s)
	for i := 0; i < 4; i++ {
		g[0][i], g[1][i] = g[1][i], g[0][i]
	}
	return get()
}

// B 将最右边一列插入到最左边
func B(s string) string {
	set(s)
	v1 := g[0][3]
	v2 := g[1][3]
	for i := 3; i >= 1; i-- {
		g[0][i], g[1][i] = g[0][i-1], g[1][i-1]
	}
	g[0][0] = v1
	g[1][0] = v2
	return get()
}
// C 将g中间四个数顺时针旋转
func C(s string) string {
	set(s)
	v1 := g[0][1]
	g[0][1] = g[1][1]
	g[1][1] = g[1][2]
	g[1][2] = g[0][2]
	g[0][2] = v1
	return get()
}

func bfs() {
	// 从start状态开始枚举,知道等于end状态
	hh, tt := 0, -1
	tt++
	q[tt] = start
	m[start] = newFather("", 0, "")
	for hh <= tt {
		head := q[hh]
		hh++
		ss := make([]string, 3)
		ss[0] = A(head)
		ss[1] = B(head)
		ss[2] = C(head)
		for i:=0;i<3;i++{
			if _, ok := m[ss[i]];!ok{
				m[ss[i]] = newFather(string(rune('A'+i)), m[head].dist + 1,head)
				tt++
				q[tt] = ss[i]
				if ss[i] == end{
					return
				}
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	tokens := strings.Split(line, " ")
	end = strings.Join(tokens, "")
	bfs()
	fmt.Println(m[end].dist)
	res := ""
	for end != start {
		res += m[end].op
		end = m[end].src
	}
	for i:=len(res) - 1;i>=0;i--{
		fmt.Print(string(res[i]))
	}
}
