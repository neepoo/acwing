package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sToSlice(a string) []int {
	l := len(a)
	res := make([]int, l)
	i := 0
	for j := l - 1; j >= 0; j-- {
		tmp, _ := strconv.Atoi(string(a[j]))
		res[i] = tmp
		i++
	}
	return res
}

// https://www.acwing.com/problem/content/793/
func Add(a, b string) string {
	res := make([]int, 0)
	numA, numB := sToSlice(a), sToSlice(b)
	c := 0 // 进喂
	for i := 0; i < len(numA) || i < len(numB); i++ {
		if i < len(numA) {
			c += numA[i]
		}
		if i < len(numB) {
			c += numB[i]
		}
		res = append(res, c%10)
		c /= 10
	}
	if c > 0 {
		res = append(res, c)
	}
	s := strings.Builder{}
	for i := len(res) - 1; i >= 0; i-- {
		s.WriteString(strconv.Itoa(res[i]))
	}
	return s.String()
}

var a, b string

func main() {
	fmt.Fscanln(os.Stdin, &a)
	fmt.Fscanln(os.Stdin, &b)
	res := Add(a, b)
	fmt.Println(res)
}
