/*
https://www.acwing.com/problem/content/872/
输出所有的约数
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var n int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &x)
		divide(x)
	}
}

func divide(x int) {
	var nums []int
	for i := 1; i <= x/i; i++ {
		if x%i == 0 {
			if x/i != i {
				nums = append(nums, i, x/i)
			}else {
				nums = append(nums, i)
			}
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for _, i := range nums {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
