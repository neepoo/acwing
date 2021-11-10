package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	t, n    int
	nums [N]int
	f [N][2]int
)



func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 500000)
	fmt.Fscanf(reader, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(reader, "%d\n", &n)
		for j := 1; j <= n; j++ {
			fmt.Fscanf(reader, "%d ", &nums[j])
		}
		for i := 1; i <= n ; i++{
			// 第i家店铺不抢的情况
			f[i][0] = max(f[i-1][0], f[i-1][1])
			// 抢第i家店铺, 这里不需要和f[i][1]取max, 因为到当前状态只有一条路线,max(a) = a, 但是这里取max会得到上一轮的状态
			f[i][1] = f[i-1][0] + nums[i]
		}
		fmt.Println(max(f[n][0], f[n][1]))
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
