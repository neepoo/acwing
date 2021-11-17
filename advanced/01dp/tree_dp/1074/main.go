/*
https://www.acwing.com/problem/content/1074/
树的最长路径
```
给定一棵树，树中包含 n 个结点（编号1~n）和 n−1 条无向边，每条边都有一个权值。

现在请你找到树中的一条最长路径。

换句话说，要找到一条路径，使得使得路径两端的点的距离最远。

注意：路径中可以只包含一个点。

输入格式
第一行包含整数 n。

接下来 n−1 行，每行包含三个整数 ai,bi,ci，表示点 ai 和 bi 之间存在一条权值为 ci 的边。

输出格式
输出一个整数，表示树的最长路径的长度。

数据范围
1≤n≤10000,
1≤ai,bi≤n,
−105≤ci≤105
输入样例：
6
5 1 6
1 4 5
6 3 9
2 6 8
6 1 7
输出样例：
22
```
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 10010
	M = 20010
)

var (
	n int

	h        [N]int
	idx      int
	e, w, ne [M]int

	ans int
)

func init() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
}

func add(a int, b int, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var a, b, c int
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		add(a, b, c)
		add(b, a, c)
	}
	dfs(1, -1)
	fmt.Println(ans)
}

func dfs(u int, p int) int {
	dist := 0 // 从该点下去的最长距离
	/*
	d1,d2表示往下走的最长距离和次长距离
	因此最终的答案就是ans = max(ans, d1+d2)
	d1, d2 可能<0, 这种情况下并不会更新d1, d2(你是负数,我不走就行了)
	d1, d2都大于0的话,说明u是最长直径候选路径中的一个点
	d1, d2只有一个大于0的话,那最长路径一定是以u为端点
	d1, d2 都不>0那我不走就可以了
	 */
	d1, d2 :=0, 0
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		v := w[i]
		// 避免套娃
		if j == p{
			continue
		}
		d := dfs(j, u) + v // 路径长度
		dist = max(dist, d)
		if d > d1{
			d2 = d1
			d1 = d
		}else if d > d2{
			d2 = d
		}
	}
	ans = max(ans, d1+d2)
	return dist
}
