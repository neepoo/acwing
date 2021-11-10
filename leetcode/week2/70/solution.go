package leetcode


func climbStairs(n int) int {
	a, b := 1, 1
	for i:=1;i<n;i++{
		c := a+b
		a = b
		b = c
	}
	return b
}