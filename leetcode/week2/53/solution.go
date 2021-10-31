package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	last, res := -0x3f3f3f3f, -0x3f3f3f3f
	for i := 0; i < len(nums); i++ {
		last = max(0, last) + nums[i]
		res = max(res, last)
	}
	return res
}
