package main

/*
https://www.acwing.com/problem/content/11/

有 N 件物品和一个容量是 V 的背包。每件物品只能使用一次。

第 i 件物品的体积是 vi，价值是 wi。

求解将哪些物品装入背包，可使这些物品的总体积不超过背包容量，且总价值最大。

输出 最优选法的方案数。注意答案可能很大，请输出答案模 109+7 的结果。

输入格式
第一行两个整数，N，V，用空格隔开，分别表示物品数量和背包容积。

接下来有 N 行，每行两个整数 vi,wi，用空格隔开，分别表示第 i 件物品的体积和价值。

输出格式
输出一个整数，表示 方案数 模 109+7 的结果。

数据范围
0<N,V≤1000
0<vi,wi≤1000

输入样例
4 5
1 2
2 4
3 4
4 6
输出样例：
2
*/
import . "fmt"

const (
	N   = 1010
	mod = int(1e9 + 7)
)

var (
	n, m int
	f, g [N]int
)

func init() {
	for i := 0; i < N; i++ {
		f[i] = -0x3f3f3f3f
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	f[0] = 0
	g[0] = 1
	Scan(&n, &m)
	var a, b int
	for i := 0; i < n; i++ {
		Scan(&a, &b)
		var maxV int
		for j := m; j >= a; j-- {
			maxV = max(f[j], f[j-a]+b)
			cnt := 0
			// 不选j得到最大价值
			if maxV == f[j] {
				cnt += g[j]
			}
			// 选j得到最大价值
			if maxV == f[j-a]+b {
				cnt += g[j-a]
			}
			g[j] = cnt % mod
			f[j] = maxV
		}
	}

	var res int
	for i := 0; i <= m; i++ {
		res = max(res, f[i])
	}
	var cnt int
	for i := 0; i <= m; i++ {
		if res == f[i] {
			cnt = (cnt + g[i]) % mod
		}
	}
	Println(cnt)
}
