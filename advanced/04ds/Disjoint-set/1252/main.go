//go:build go1.16
// +build go1.16

/*
Alice和Bob玩了一个古老的游戏：首先画一个 n×n 的点阵（下图 n=3 ）。

接着，他们两个轮流在相邻的点之间画上红边和蓝边：

1.png

直到围成一个封闭的圈（面积不必为 1）为止，“封圈”的那个人就是赢家。因为棋盘实在是太大了，他们的游戏实在是太长了！

他们甚至在游戏中都不知道谁赢得了游戏。

于是请你写一个程序，帮助他们计算他们是否结束了游戏？

输入格式
输入数据第一行为两个整数 n 和 m。n表示点阵的大小，m 表示一共画了 m 条线。

以后 m 行，每行首先有两个数字 (x,y)，代表了画线的起点坐标，接着用空格隔开一个字符，假如字符是 D，则是向下连一条边，如果是 R 就是向右连一条边。

输入数据不会有重复的边且保证正确。

输出格式
输出一行：在第几步的时候结束。

假如 m 步之后也没有结束，则输出一行“draw”。

数据范围
1≤n≤200，
1≤m≤24000
输入样例：
3 5
1 1 D
1 1 R
1 2 D
2 1 R
2 2 D
输出样例：
4

两个点如果属于同一个集合那么说明游戏结束了
*/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 40010 * 2 // 边的上限 200*200 + 10
var (
	n, m   int // 方格的大小以及边的数量
	x, y   int // 坐标(读取的时候从1开始,实际使用的时候需要--)
	op     string
	ans    int // 0表示平局
	p      [N]int
	pa, pb int
)

// get 将二维坐标唯一映射成一维的点(x, y >=0)
func get(x, y int) int {
	return n*x + y
}

// find 给出根节点的编号
func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func init() {
	for i := 0; i < N; i++ {
		p[i] = i
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscan(reader, &n, &m)
	for i := 1; i <= m; i++ {
		Fscan(reader, &x, &y, &op)
		x--
		y--
		switch op {
		case "D":
			pa, pb = find(get(x, y)), find(get(x+1, y))
		default:
			pa, pb = find(get(x, y)), find(get(x, y+1))
		}
		if pa == pb {
			ans = i
			break
		}
		p[pa] = pb
	}
	if ans > 0 {
		Println(ans)
	} else {
		Println("draw")
	}
}
