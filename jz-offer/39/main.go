package main

func majorityElement(nums []int) int {
	//m := make(map[int]int)
	n := len(nums)
	//for _, v := range nums{
	//	m[v]++
	//	if m[v] >= n/2 +1{
	//		return v
	//	}
	//}
	//return -1
	var (
		res = nums[0]
		cnt =1
	)
	for i := 1; i < n; i++ {
		if nums[i] == res{
			cnt++
		}else {
			cnt--
			if cnt == 0{
				res = nums[i+1]
			}
		}
	}
	return res
}


func main() {
	res := majorityElement([]int{2,2,1,1,1,2,2})
	println(res)
}
