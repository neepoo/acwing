//https://www.acwing.com/problem/content/799/
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
	n, m    int
	l, r, c int
	a       [N]int
)

func add(l, r, c int) {
	a[l] += c
	a[r+1] -= c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ss := strings.Builder{}
	fmt.Fscan(reader, &n, &m)
	var x int
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &x)
		add(i, i, x)
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &l, &r, &c)
		add(l, r, c)
	}
	for i := 1; i <= n; i++ {
		a[i] += a[i-1]
		if i < n {
			ss.WriteString(strconv.Itoa(a[i]) + " ")
		} else {
			ss.WriteString(strconv.Itoa(a[i]))
		}
	}
	fmt.Println(ss.String())
}
