package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1e6 + 10

var (
	n int
	q [N]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	Fscan(reader, &n)
	for i := 0; i < n; i++ {
		Fscan(reader, &q[i])
	}
	mergeSort(q[:n], 0, n-1)
	ss := strings.Builder{}
	for i := 0; i < n; i++ {
		ss.WriteString(strconv.Itoa(q[i]))
		if i < n-1 {
			ss.WriteString(" ")
		}
	}
	Println(ss.String())
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)

	tmp := make([]int, r-l+1)
	i, j := l, mid+1
	k := 0
	for i <= mid && j <= r {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j <= r {
		tmp[k] = nums[j]
		j++
		k++
	}
	copy(nums[l:r+1], tmp)
}
