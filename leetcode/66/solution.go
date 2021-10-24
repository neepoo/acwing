package solution

func plusOne(digits []int) []int {
	c := 0
	res := make([]int, len(digits)+1)
	tmp := -1
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			tmp = digits[i] + c + 1
		} else {
			tmp = digits[i] + c
		}
		c = tmp / 10
		res[i+1] = tmp % 10
	}
	if c == 1 {
		res[0] = 1
		return res
	} else {
		return res[1:]
	}
}

func plusOneNew(digits []int) []int {
	c := 1
	res := make([]int, len(digits)+1)
	tmp := -1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp = digits[i] + c
		c = tmp / 10
		res[i+1] = tmp % 10
	}
	if c == 1 {
		res[0] = 1
		return res
	} else {
		return res[1:]
	}
}