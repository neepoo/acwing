package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	var f [12]int
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1]
		t := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
		if t >= 10 && t <= 25 {
			f[i] += f[i-2]
		}
	}
	return f[n]
}
func main() {
	//fmt.Println(translateNum(12258))
	fmt.Println(translateNum(12))
}
