/*
https://www.acwing.com/problem/content/1254/
Joe觉得云朵很美，决定去山上的商店买一些云朵。

商店里有 n 朵云，云朵被编号为 1,2,…,n，并且每朵云都有一个价值。

但是商店老板跟他说，一些云朵要搭配来买才好，所以买一朵云则与这朵云有搭配的云都要买。

但是Joe的钱有限，所以他希望买的价值越多越好。

输入格式
第 1 行包含三个整数 n，m，w，表示有 n 朵云，m 个搭配，Joe有 w 的钱。

第 2∼n+1行，每行两个整数 ci，di 表示 i 朵云的价钱和价值。

第 n+2∼n+1+m 行，每行两个整数 ui，vi，表示买 ui 就必须买 vi，同理，如果买 vi 就必须买 ui。

输出格式
一行，表示可以获得的最大价值。

数据范围
1≤n≤10000,
0≤m≤5000,
1≤w≤10000,
1≤ci≤5000,
1≤di≤100,
1≤ui,vi≤n
输入样例：
5 3 10
3 10
3 10
3 10
5 100
10 1
1 3
3 2
4 2
输出样例：
1
*/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 10010

type pair struct {
	root, money, value int
}

var (
	n, m, w int
	q       [10010]pair
	f       [N]int
)

func find(x int) int {
	if q[x].root != x {
		q[x].root = find(q[x].root)
	}
	return q[x].root
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscan(reader, &n, &m, &w)
	var money, value int
	for i := 1; i <= n; i++ {
		Fscan(reader, &money, &value)
		q[i] = pair{i, money, value}
	}
	var a, b int
	for i := 0; i < m; i++ {
		Fscan(reader, &a, &b)
		pa, pb := find(a), find(b)
		if pa != pb {
			q[pa].root = pb
			q[pb].money += q[pa].money
			q[pb].value += q[pa].value
		}
	}

	var cnt = 0
	for i := 1; i <= n; i++ {
		if q[i].root == i {
			cnt++
			q[cnt] = q[i]
		}
	}
	for i := 1; i <= cnt; i++ {
		for j := w; j >= q[i].money; j-- {
			f[j] = max(f[j], f[j-q[i].money]+q[i].value)
		}
	}
	Println(f[w])
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2

}
