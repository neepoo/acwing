package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	var res [][]int
	i, j, sum := 1, 1, 1
	for ; i <= target/2; i++ {
		for sum < target {
			j++
			sum += j
		}
		if sum == target && j > i {
			tmp := make([]int, j-i+1)
			for k := 0; k <= j-i; k++ {
				tmp[k] = i + k
			}
			res = append(res, tmp)
		}
		sum -= i
	}
	return res
}
func main() {
	fmt.Println(findContinuousSequence(9))
}
