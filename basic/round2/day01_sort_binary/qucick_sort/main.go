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
	quickSort(q[:n], 0, n-1)
	ss := strings.Builder{}
	for i := 0; i < n; i++ {
		ss.WriteString(strconv.Itoa(q[i]))
		if i < n-1 {
			ss.WriteString(" ")
		}
	}
	Println(ss.String())
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	x := nums[(l+r)/2]
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
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}
