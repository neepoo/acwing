package main
// 一次买入卖出记做一次交易
// 买入时交易数量+1
// 卖出时交易数量保持不变
import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 100010
	M   = 110
	INF = 0x3f3f3f3f
)

var (
	nums [N]int
	f    [N][M][2]int
	n, k int
)

func init() {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			for k := 0; k < 2; k++ {
				f[i][j][k] = -INF
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	for i := 0; i < n; i++ {
		f[i][0][0] = 0 // 第i天,进行了0笔交易,当前手上没股票的最大收益
		fmt.Fscanf(reader, "%d ", &nums[i+1])
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			//  第i-1天,已经进行了j笔交易(包括在途),并且手上有股票            第i-1天,已经进行了j-1笔交易,并且手上没有股票
			// 	|									   			/
			//  |								    	/
			//  |								 	/
			//  ↧							  ↙
			// 第i天,已经进行了j笔交易,并且手上有股票
			f[i][j][1] = max(f[i-1][j][1], f[i-1][j-1][0]-nums[i])
			//  第i-1天,已经进行了j笔交易,并且手上没有股票
			f[i][j][0] = max(f[i-1][j][0]/*继续观望,不出手,不买*/, f[i-1][j][1]+nums[i]/*第i-1天还持有股票, 第i天卖了,获取收益nums[i]*/)
		}
	}
	var res int
	for i := 1; i <= k; i++ {
		res = max(res, f[n][i][0])
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
