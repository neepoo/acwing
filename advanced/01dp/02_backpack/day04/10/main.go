package main

/*
https://www.acwing.com/problem/content/10/

有 N 个物品和一个容量是 V 的背包。

物品之间具有依赖关系，且依赖关系组成一棵树的形状。如果选择一个物品，则必须选择它的父节点。

如下图所示：


如果选择物品5，则必须选择物品1和2。这是因为2是5的父节点，1是2的父节点。

每件物品的编号是 i，体积是 vi，价值是 wi，依赖的父节点编号是 pi。物品的下标范围是 1…N。

求解将哪些物品装入背包，可使物品总体积不超过背包容量，且总价值最大。

输出最大价值。

输入格式
第一行有两个整数 N，V，用空格隔开，分别表示物品个数和背包容量。

接下来有 N 行数据，每行数据表示一个物品。
第 i 行有三个整数 vi,wi,pi，用空格隔开，分别表示物品的体积、价值和依赖的物品编号。
如果 pi=−1，表示根节点。 数据保证所有物品构成一棵树。

输出格式
输出一个整数，表示最大价值。

数据范围
1≤N,V≤100
1≤vi,wi≤100
父节点编号范围：

内部结点：1≤pi≤N;
根节点 pi=−1;
输入样例
5 7
2 3 -1
2 2 1
3 5 1
4 7 2
3 6 2
输出样例：
11
难度：困难
时/空限制：1s / 64MB
总通过数：6326
总尝试数：10459
来源：背包九讲
算法标签
*/
import . "fmt"

const N = 110

var (
	n, m           int
	h, e, ne, v, w [N]int
	idx, root      int
	f              [N][N]int
)

func init() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func dfs(u int) {
	for i := h[u]; i != -1; i = ne[i] {
		nid := e[i] // 子节点编号
		dfs(nid)
		// 对以u为根节点的子树做分组背包
		for j := m - v[u]; j >= 0; j-- { // 枚举体积
			for k := 0; k <= j; k++ {
				// 						空出k给它的儿子
				f[u][j] = max(f[u][j], f[u][j-k]+f[nid][k])
			}
		}
	}
	for j := m; j >= v[u]; j-- {
		f[u][j] = f[u][j-v[u]] + w[u]
	}
	for j := 0; j < v[u]; j++ {
		f[u][j] = 0
	}
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

func main() {
	Scan(&n, &m)
	var p int
	for i := 1; i <= n; i++ {
		Scan(&v[i], &w[i], &p)
		if p == -1 {
			root = i // 根节点
		} else {
			add(p, i)
		}
	}
	dfs(root)
	Println(f[root][m])
}
