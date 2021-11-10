package day02


func duplicateInArray(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums{
		m[v]++
		if m[v] > 1{
			return v
		}
	}
	return -1
}