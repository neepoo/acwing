package day03

func searchArray(nums [][]int, target int) bool {
	row := len(nums)
	if row == 0{
		return false
	}
	col := len(nums[0])
	if col == 0{
		return false
	}
	i, j := 0, col-1
	for i < row && j >= 0{
		v:= nums[i][j]
		if v == target{
			return true
		}else if v > target{
			j --
		}else {
			i ++
		}
	}
	return false
}
