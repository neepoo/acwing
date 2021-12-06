package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	x, y := 0, 0
	// 找右区间
	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if nums[l] != target {
		return 0
	}
	x = l
	l, r = 0, n-1
	// 找左区间
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	y = r
	return x - y + 1
}

func missingNumber(nums []int) int {
	n := len(nums)
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		switch {
		case nums[mid] < mid:
			l = mid
		case nums[mid] > mid:
			r = mid
		default:
			l = mid + 1
		}
	}
	return l
}

func main() {
	//fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(missingNumber([]int{0, 1, 3}))
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
	fmt.Println(missingNumber([]int{0}))
}
