package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return getReverse(x) == x
}

func getReverse(x int) int {
	var res = 0
	for x > 0{
		res = res * 10 + x%10
		x/=10
	}
	return res
}


