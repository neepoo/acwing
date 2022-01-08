package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1010

var (
	n, m, q           int
	b                 [N][N]int
	x1, y1, x2, y2, c int
)

func add(x1, y1, x2, y2, c int) {
	b[x1][y1] += c
	b[x2+1][y1] -= c
	b[x1][y2+1] -= c
	b[x2+1][y2+1] += c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, N*N*4)
	Fscanf(reader, "%d %d %d\n", &n, &m, &q)
	//Println("n m q", n, m, q)
	var t int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			Fscanf(reader, "%d ", &t)
			//Println("t", t)
			add(i, j, i, j, t)
		}
	}
	for i := 0; i < q; i++ {
		Fscanf(reader, "%d %d %d %d %d\n", &x1, &y1, &x2, &y2, &c)
		//Println(x1, y1, x2, y2, c)
		add(x1, y1, x2, y2, c)
	}
	res := strings.Builder{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			b[i][j] += b[i][j-1] + b[i-1][j] - b[i-1][j-1]
			res.WriteString(strconv.Itoa(b[i][j]) + " ")
		}
		res.WriteString("\n")
	}
	Print(res.String())
}
