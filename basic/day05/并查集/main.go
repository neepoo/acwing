/*
 应用
	1. 合并两个集合
	2. 询问两个元素是否属于同一个集合
每个集合用一颗树来维护,每个集合的根节点的编号就是当前集合 bianhao

q1: is root? x?= p(x)
q2: find root : while(p(x)!=x) x= p[x]
q3: how merge two tree?
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100010

var (
	n, m int
	p    [N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 5*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() && m > 0 {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		op := tokens[0]
		a, _ := strconv.Atoi(tokens[1])
		b, _ := strconv.Atoi(tokens[2])
		switch op {
		case "M":
			p[find(a)] = find(b)
		case "Q":
			if find(a) == find(b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
		m--
	}
}

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}
