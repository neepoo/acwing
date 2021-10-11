package leetcode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var cnt int // val cnt
	for i := 0; i < len(nums); i++ {
		if nums[i] != val{
			nums[cnt] = nums[i]
			cnt++
		}
	}
	return cnt
}
