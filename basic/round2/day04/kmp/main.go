//https://www.acwing.com/problem/content/833/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1e5 + 10

func main() {
	var (
		n, m int
		p, s string
		next [N]int
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &p, &m, &s)
	p = " " + p
	s = " " + s
	// 求next, next[1] = 0
	j := 0
	for i := 2; i <= n; i++ {
		for j > 0 && p[i] != p[j+1] {
			j = next[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		next[i] = j
	}
	j = 0
	ss := strings.Builder{}
	for i := 1; i <= m; i++ {
		for j > 0 && s[i] != p[j+1] {
			j = next[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == n {
			ss.WriteString(strconv.Itoa(i-n) + " ")
			j = next[j]
		}
	}
	fmt.Println(ss.String())
}

/*
3
aba
5
ababa

want:
	0 2
*/
