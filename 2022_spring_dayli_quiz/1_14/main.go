/*
https://www.acwing.com/problem/content/1954/
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	x, y int
}

const N = 40010

var (
	n, x, y, z int
	ps         []pair
	s          [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &x, &y, &z)
	ps = make([]pair, n)
	nums := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ps[i].x, &ps[i].y)
		nums = append(nums, ps[i].x, ps[i].y)
	}
	sort.Ints(nums)
	nums = unique(nums)
	for _, p := range ps {
		l, r := find(nums, p.x), find(nums, p.y)
		s[l] += y
		s[r+1] -= y
		s[1] += x
		s[l] -= x
		s[r+1] += z
	}
	var res int
	for i := 1; i <= len(nums); i++ {
		s[i] += s[i-1]
		res = max(res, s[i])
	}
	fmt.Println(res)
}

func unique(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	cnt := 1
	last := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] != last {
			nums[cnt] = nums[i]
			cnt++
			last = nums[i]
		}
	}
	return nums[:cnt]
}

func find(nums []int, x int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r + 1
}
