// 多重背包的二进制优化方法。
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 25000

var (
	f    [N]int
	v, w [N]int
	n, m int
)

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 4*N)
	Fscanf(reader, "%d %d\n", &n, &m)
	cnt := 0 // 新的物品数量
	var a, b, c int
	for i := 1; i <= n; i++ {
		Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		k := 1
		for k <= c {
			cnt++
			v[cnt] = a * k
			w[cnt] = b * k
			c -= k
			k *= 2
		}
		if c > 0 {
			cnt++
			v[cnt] = a * c
			w[cnt] = b * c
		}
	}
	n = cnt
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}
	Println(f[m])
}
