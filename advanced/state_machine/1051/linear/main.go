package main
// https://www.acwing.com/problem/content/description/1051/
// 3460ms
import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	t, n    int
	f, nums [N]int
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
		f[1] = nums[1]
		for i := 2; i <= n; i++ {
			f[i] = max(f[i-1], max(f[i], f[i-2]+nums[i]))
		}
		fmt.Println(f[n])
		for i := 0; i < N; i++ {
			f[i] = 0
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
