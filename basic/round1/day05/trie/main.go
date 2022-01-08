// trie 高效的存储和查找字符串集合的数据结构
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const N = 100010

var (
	son [N][26]int
	cnt [N]int
	idx int
)

func sub(a, b rune) int32 {
	return a - b
}

func insert(s string) {
	p := 0 // p代表每个节点的编号
	for _, i := range []rune(s) {
		u := sub(i, 'a')
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
	cnt[p]++
}

func find(s string) int {
	p := 0
	for _, i := range []rune(s) {
		u := sub(i, 'a')
		if son[p][u] == 0 {
			return 0
		}
		p = son[p][u]
	}
	return cnt[p]
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*6*10000)
	fmt.Fscanf(reader, "%d\n", &n)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() && n > 0 {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		op := tokens[0]
		s := tokens[1]
		switch op {
		case "I":
			insert(s)
		case "Q":
			fmt.Println(find(s))
		}
		n--
	}
}
