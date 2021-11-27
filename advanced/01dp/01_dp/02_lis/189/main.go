/*
https://www.acwing.com/problem/content/189/

为了对抗附近恶意国家的威胁，R 国更新了他们的导弹防御系统。

一套防御系统的导弹拦截高度要么一直 严格单调 上升要么一直 严格单调 下降。

例如，一套系统先后拦截了高度为 3 和高度为 4 的两发导弹，那么接下来该系统就只能拦截高度大于 4 的导弹。

给定即将袭来的一系列导弹的高度，请你求出至少需要多少套防御系统，就可以将它们全部击落。

输入格式
输入包含多组测试用例。

对于每个测试用例，第一行包含整数 n，表示来袭导弹数量。

第二行包含 n 个不同的整数，表示每个导弹的高度。

当输入测试用例 n=0 时，表示输入终止，且该用例无需处理。

输出格式
对于每个测试用例，输出一个占据一行的整数，表示所需的防御系统数量。

数据范围
1≤n≤50
输入样例：
5
3 5 2 4 1
0
输出样例：
2
样例解释
对于给出样例，最少需要两套防御系统。

一套击落高度为 3,4 的导弹，另一套击落高度为 5,2,1 的导弹。
*/
package main

import (
	"bufio"
	. "fmt"
	"os"
)

const N = 55

var (
	n, res      int
	a, up, down [N]int
)

func dfs(u, uc, dc int) {
	if uc+dc >= res {
		return
	}
	if u == n {
		res = uc + dc
		return
	}

	// 先尝试把第u枚导弹放到上升系统中
	k := 0
	for k < uc && a[u] <= up[k] {
		k++
	}
	t := up[k]
	up[k] = a[u]
	if k >= uc {
		dfs(u+1, uc+1, dc)
	} else {
		dfs(u+1, uc, dc)
	}
	up[k] = t

	// 尝试放到下降系统中
	k = 0
	for k < dc && a[u] >= down[k] {
		k++
	}
	t = down[k]
	down[k] = a[u]
	if k >= dc {
		dfs(u+1, uc, dc+1)
	} else {
		dfs(u+1, uc, dc)
	}
	down[k] = t
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		Fscan(reader, &n)
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			Fscan(reader, &a[i])
		}
		res = n // 最坏情况,一枚导弹一套系统
		dfs(0, 0, 0)
		Println(res)
	}
}
