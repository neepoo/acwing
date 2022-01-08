// 利用归并排序求解数组的逆序对个数

package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	n    int
	nums [N]int
)

func mergeSort(l, r int) uint64 {
	if l >= r {
		return 0
	}
	tmp := make([]int, r-l+1)
	mid := (l + r) / 2
	//     左半边的的逆序对数量    右半边的逆序对数量
	res := mergeSort(l, mid) + mergeSort(mid+1, r)
	i, j, k := l, mid+1, 0
	for i <= mid && j <= r {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			// 左边的数>右边的数,也会产生逆序对
			res += uint64(mid - i + 1)
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = nums[j]
		k++
		j++
	}
	copy(nums[l:r+1], tmp)
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader = bufio.NewReaderSize(reader, 2*N)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &nums[i])
	}
	fmt.Println(mergeSort(0, n-1))
}
