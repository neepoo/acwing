package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	n, k int
	nums [N]int
)

func quickSelect(l, r, k int) int {
	if l >= r {
		// 区间只有一个数
		return nums[l]
	}
	x := nums[(l+r)/2]
	i, j := l-1, r+1
	/*
	这个数据说明了为什么 i++,j--要写在循环内部
	10 10
	50 43 63 97 30 89 89 94 30 33
	*/
	for i < j {
		i++
		for nums[i] < x {
			i++
		}
		j--
		for nums[j] > x {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	sl := j - l + 1 // 左边有这么多个数
	if k <= sl {
		// 第k小的数在左边
		return quickSelect(l, j, k)
	}
	// 第k小的数是右边第k-sl小的数
	return quickSelect(j+1, r, k-sl)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &nums[i])
	}
	fmt.Println(quickSelect(0, n-1, k))
}
