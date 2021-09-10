package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.acwing.com/problem/content/794/

var a, b string

func main() {
	fmt.Fscanln(os.Stdin, &a)
	fmt.Fscanln(os.Stdin, &b)

	fmt.Println(sub(a, b))
}

// 判断 a >= b?
func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	} else {
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] != b[i] {
				return a[i] > b[i]
			}
		}
		return true
	}
}

// 需要处理前导0
func removePrefixZero(a []int) []int{
	for len(a) > 1 && a[len(a) -1] == 0{
		a = a[:len(a) - 1]
	}
	return a
}

// 保证a >= b
func subSlice(a, b []int) string {
	var c []int
	res := strings.Builder{}
	t := 0
	for i := 0; i < len(a); i++ {
		t = a[i] - t
		if i < len(b) {
			t -= b[i]
		}
		c = append(c, (t+10)%10)
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}
	// 处理前导0
	c = removePrefixZero(c)
	for i := len(c) - 1; i >= 0; i-- {
		res.WriteString(strconv.Itoa(c[i]))
	}
	return res.String()
}

// 需要保证a表示的数>b表示的数
func sub(a, b string) string {
	numA, numB := sToSlice(a), sToSlice(b)
	if cmp(numA, numB) {
		return subSlice(numA, numB)
	} else {
		return "-" + subSlice(numB, numA)
	}
}

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
