package main

import (
	"bufio"
	"fmt"
	"os"
)

// 二维费用背包问题, 一维是精灵球的数量,另外一维是皮卡丘的体力
// 知道套路之后,直接套01背包模板即可

const (
	N = 1010 //  精灵球的数量上限
	H = 510  // 皮卡丘的体力上限

)

var (
	f         [N][H]int //能收复精灵数量的最大值
	n, h, cnt int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d\n", &n, &h, &cnt)
	for i := 0; i < cnt; i++ {
		var n1, h1 int
		fmt.Fscanf(reader, "%d %d\n", &n1, &h1)

		for j := n; j >= n1; j-- {
			for k := h - 1; k >= h1; k-- {
				f[j][k] = max(f[j][k], f[j-n1][k-h1]+1)
			}
		}

	}
	fmt.Printf("%d ", f[n][h-1])
	k := h - 1 // 消耗的体力
	for k > 0 && f[n][k-1] == f[n][h-1] {
		k--
	}
	fmt.Println(h - k /*剩余的最大体力*/)
}
