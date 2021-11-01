package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	O   = 1000  // 氧气
	N   = 1000 // 氮气
	MAX = 0x3F3F3F3F
)

var (
	o, n, k int
	f       [O][N]int
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func init() {
	for i := 0; i < O; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = MAX
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 1000000)
	fmt.Fscanf(reader, "%d %d\n", &o, &n)
	fmt.Fscanf(reader, "%d\n", &k)
	// 恰好选v1氧气和v2氮气
	f[0][0] = 0 // 氧气不要,氮气不要,因此体积是0
	for k > 0{
		var a, b, c int
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		for i:=O-1; i>=a; i--{
			for j:=N-1; j >= b; j--{
				f[i][j] = min(f[i][j], f[i-a][j-b] + c)
			}
		}
		k --
	}
	ans := int(1e9)
	for i := o; i < O; i++ {
		for j := n; j < N; j++ {
			ans = min(ans, f[i][j])
		}
	}
	fmt.Println(ans)
}
