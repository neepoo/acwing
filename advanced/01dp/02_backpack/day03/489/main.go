package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 32010 // 前的总数
	M = 70    // 物品总数
)

type line struct {
	a, b, c int
}


var (
	v    [M][]int // [][0]表示选主
	w    [M][]int
	n, m int
	f    [N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	lines := make([]line, m)
	var a, b, c int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		lines[i] = line{a,b,c}
	}
	for i := 1; i <= m; i++ {
		a, b, c = lines[i-1].a, lines[i-1].b, lines[i-1].c
		if c == 0 {
			v[i] = append(v[i], a)
			w[i] = append(w[i], a*b)
		} else {
			v[c] = append(v[c], a + v[c][0])
			w[c] = append(w[c], a*b + w[c][0])
			if len(v[c]) == 3 {
				v[c] = append(v[c], v[c][1]+v[c][2] - v[c][0])
			}
			if len(w[c]) == 3 {
				w[c] = append(w[c], w[c][1]+w[c][2] - w[c][0])
			}
		}
	}
	for i := 1; i <= m; i++ {
		for j := n; j >= 0; j-- {
			for k := 0; k < len(v[i]); k++ {
				if j >= v[i][k] {
					f[j] = max(f[j], f[j-v[i][k]]+w[i][k])
				}
			}
		}
	}
	fmt.Println(f[n])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcm(a, b int) int{
	if b == 0{
		return a
	}
	return gcm(b, a%b)
}