package dichotomous

func numRange(nums []int, x int) (int, int) {
	i := 0
	j := len(nums) - 1
	var left int
	var right int
	// 找左边界 (右半边的数都>=x)
	for i < j {
		mid := (i + j) / 2
		if nums[mid] >= x {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if x != nums[i] {
		return -1, -1
	}
	left = i

	i = 0
	j = len(nums) - 1
	// 找右边界(左边的数都<=x)
	for i < j {
		mid := (i + j + 1) / 2
		if nums[mid] <= x {
			i = mid
		} else {
			j = mid - 1
		}
	}
	right = j
	return left, right

}
