package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	var res string
	sort.Slice(nums, func(i, j int) bool {
		s1 := strconv.Itoa(nums[i])
		s2 := strconv.Itoa(nums[j])
		return s1+s2 < s2+s1
	})
	for _, v := range nums {
		res += strconv.Itoa(v)
	}
	return res
}

func main() {
	minNumber([]int{3, 30, 34, 5, 9})
	fmt.Println("330" < "303")
}
