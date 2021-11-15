package main

func minArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	n := len(nums) - 1
	for n > 0 && nums[n] == nums[0] {
		n--
	}
	if nums[n] >= nums[0] {
		return nums[0]
	}
	l, r := 0, n
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < nums[0] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
func main() {
	println(minArray([]int{2, 2, 2, 0, 1}))
}
