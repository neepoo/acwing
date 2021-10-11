package leetcode

func removeDuplicates(nums []int) int {
	res := 1
	if len(nums) == 0 {
		return 0
	}
	lastElem := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != lastElem {
			lastElem = nums[i]
			nums[res] = lastElem
			res++
		}
	}

	return res
}
