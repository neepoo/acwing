package main

func reOrderArray(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		for nums[i]%2 != 0 {
			i++
		}
		for nums[j]%2 != 1 {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
