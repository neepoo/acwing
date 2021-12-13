/*
https://www.acwing.com/problem/content/240/
有一个划分为 N 列的星际战场，各列依次编号为 1,2,…,N。

有 N 艘战舰，也依次编号为 1,2,…,N，其中第 i 号战舰处于第 i 列。

有 T 条指令，每条指令格式为以下两种之一：

M i j，表示让第 i 号战舰所在列的全部战舰保持原有顺序，接在第 j 号战舰所在列的尾部。
C i j，表示询问第 i 号战舰与第 j 号战舰当前是否处于同一列中，如果在同一列中，它们之间间隔了多少艘战舰。
现在需要你编写一个程序，处理一系列的指令。

输入格式
第一行包含整数 T，表示共有 T 条指令。

接下来 T 行，每行一个指令，指令有两种形式：M i j 或 C i j。

其中 M 和 C 为大写字母表示指令类型，i 和 j 为整数，表示指令涉及的战舰编号。

输出格式
你的程序应当依次对输入的每一条指令进行分析和处理：

如果是 M i j 形式，则表示舰队排列发生了变化，你的程序要注意到这一点，但是不要输出任何信息；

如果是 C i j 形式，你的程序要输出一行，仅包含一个整数，表示在同一列上，第 i 号战舰与第 j 号战舰之间布置的战舰数目，如果第 i 号战舰与第 j 号战舰当前不在同一列上，则输出 −1。

数据范围
N≤30000,T≤500000
输入样例：
4
M 2 3
C 1 2
M 2 4
C 4 2
输出样例：
-1
1
*/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 30010

var (
	t int
	//d[x]表示x到fa[x]的距离
	fa, size, d [N]int // x -> y -> z 这种情况下,有些时候d[z]表示的是z这个节点到他的直接根节点y的距离,但是最终一定会更新成到它祖宗x的距离
)

func find(x int) int {
	if x != fa[x] {
		root := find(fa[x])
		d[x] += d[fa[x]] // 回溯的时候d[fa[x]]一定就表示fa[x]到他祖宗的距离, d[x]当前只表示到fa[x]的距离,因此x到祖宗的距离就是d[x]+=d[fa[x]]
		fa[x] = root
	}
	return fa[x]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func init() {
	for i := 1; i < N; i++ {
		fa[i] = i
		size[i] = 1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscan(reader, &t)
	var a, b int
	var op string
	for i := 0; i < t; i++ {
		Fscan(reader, &op, &a, &b)

		if op == "M" {
			pa, pb := find(a), find(b)
			fa[pa] = pb
			d[pa] = size[pb]
			size[pb] += size[pa]
		} else {
			// C
			pa, pb := find(a), find(b)
			if pa != pb {
				Println("-1")
			} else {
				b2 := abs(d[a]-d[b]) - 1
				Println(max(0, b2))
			}
		}
	}
}
