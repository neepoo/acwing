package main

import (
	"bufio"
	"os"

	. "fmt" // 借鉴0x3f大佬的做法
)

//https://www.acwing.com/problem/content/799/
const N = 100010

var (
	n, m    int
	b       [N]int
	l, r, c int
)

func add(l, r, c int) {
	b[l] += c
	b[r+1] -= c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, N*4)
	Fscanf(reader, "%d %d\n", &n, &m)
	//Println(n, m)
	var t int
	for i := 1; i <= n; i++ {
		Fscanf(reader, "%d ", &t)
		add(i, i, t)
	}
	for i := 0; i < m; i++ {
		Fscanf(reader, "%d %d %d\n", &l, &r, &c)
		add(l, r, c)
	}
	// 计算前缀和
	for i := 1; i <= n; i++ {
		b[i] += b[i-1]
		Printf("%d ", b[i])
	}
}
