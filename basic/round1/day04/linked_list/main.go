// 单链表用来表示邻接表,邻接表用来存储树和图
// https://www.acwing.com/problem/content/828/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N int = 1e5 + 10

var (
	e    [N]int
	ne   [N]int
	head int // 表示头节点的下标
	idx  int // 存储当前用到的下标
)

func init() {
	idx = 1
	head = -1
	for i := 0; i < N; i++ {
		ne[i] = -1
	}
}

func addToHead(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

// x 插入到下标为k的点后面
func add(x int, k int) {
	e[idx] = x
	ne[idx] = ne[k]
	ne[k] = idx
	idx++
}

// 将下标时k的点后面的节点删掉
func del(k int) {
	ne[k] = ne[ne[k]]
}

func put() {
	for i := head; i != -1; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
}

func main() {
	var m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &m)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() && m > 0 {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "H":
			x, _ := strconv.Atoi(tokens[1])
			addToHead(x)
		case "D":
			k, _ := strconv.Atoi(tokens[1])
			if k == 0 {
				head = ne[head]
			} else {

				del(k)
			}
		case "I":
			k, _ := strconv.Atoi(tokens[1])
			x, _ := strconv.Atoi(tokens[2])
			add(x, k)
		}
		m--
	}
	put()
}
