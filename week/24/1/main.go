package main

import "fmt"

const N = 11
var (
	n int
	nums [N]int
)
func main() {
	fmt.Scan(&n)
	var m = -0x3f3f3f3f
	for i:=0;i<n;i++{
		fmt.Scan(&nums[i])
		m = max(m, nums[i])
	}
	fmt.Println(m^nums[n-1])
}

func max(m int, i int) int {
	if m > i{
		return m
	}
	return i
}
