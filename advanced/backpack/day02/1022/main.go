package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 22 // 氧气
	M   = 80 // 氮气
	MAX = 0x3F3F3F3F
)

var (
	n, m, k int
	f       [N][M]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func init() {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			f[i][j] = MAX
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 1000000)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	fmt.Fscanf(reader, "%d\n", &k)
	// 恰好选v1氧气和v2氮气
	f[0][0] = 0 // 氧气不要,氮气不要,因此体积是0
	for k > 0{
		var v1, v2, w int
		fmt.Fscanf(reader, "%d %d %d\n", &v1, &v2, &w)
		for i:= n; i>=0; i--{
			for j:= m; j >= 0; j--{
				f[i][j] = min(f[i][j], f[max(0, i-v1)][max(0, j-v2)] +w)
			}
		}
		k --
	}
	fmt.Println(f[n][m])
}

func min(i int, i2 int) int {
	if i < i2{
		return i
	}
	return i2
}
