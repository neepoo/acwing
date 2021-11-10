//https://www.acwing.com/problem/content/167/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 20

var (
	nums [N]int
	sum  [N]int
	n, m int
	res  = N
)

// u:第几只猫[0,n), k当前可用车的数量
func dfs(u, k int) {
	// 最优性减枝
	if k >= res {
		return
	}
	if u == n {
		// 所有的猫都已经装箱了
		res = k
		return
	}

	for i := 0; i < k; i++ {
		if sum[i]+nums[u] <= m {
			sum[i] += nums[u]
			dfs(u+1, k)
			sum[i] -= nums[u]

		}
	}
	// 新开一辆车
	sum[k] = nums[u]
	// ,枚举下一层,多一辆车
	dfs(u+1, k+1)
	sum[k] = 0
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &nums[i])
	}
	sort.Slice(nums[:n], func(i, j int) bool {
		return nums[i] > nums[j]
	})
	// 从第一只猫开始装, 当前使用的缆车数量是0
	dfs(0, 0)
	fmt.Println(res)
}
