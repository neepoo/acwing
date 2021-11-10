package leetcode

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	if len(nums1) % 2 == 0{
		return (float64(nums1[(len(nums1)-1)/2]) + float64(nums1[len(nums1)/2])) / 2
	}else {
		return float64(nums1[len(nums1)/2])
	}

}