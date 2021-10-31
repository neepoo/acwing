package solution

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	mid := (l + r + 1) / 2
	for l < r {
		mid = (l + r + 1) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	//log.Printf("l=%d, r = %d, mid=%d\n", l, r, mid)
	if nums[l] == target {
		return l
	} else if nums[l] < target {
		return l + 1
	}
	if l == 0 {
		return 0
	}
	return l - 1
}

// 从左往右,找到第一个大于等于target的数
func searchInsertNew(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
