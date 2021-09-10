package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	a string
	b int
)

func main() {
	fmt.Fscanln(os.Stdin, &a)
	fmt.Fscanln(os.Stdin, &b)
	fmt.Println(Mul(a, b))
}

func Mul(a string, b int) string {
	return mul(sToSlice(a), b)
}

func mul(a []int, b int) string {
	t := 0
	var c []int
	res := strings.Builder{}
	for i := 0; i < len(a) || t > 0; i++ {
		if i < len(a) {
			t += b * a[i]
		}
		c = append(c, t%10)
		t /= 10
	}
	// 去除前导0
	for i := len(c) - 1; i >= 1 && c[i] == 0; i-- {
		c = c[:len(c)-1]
	}
	for i := len(c) - 1; i >= 0; i-- {
		res.WriteString(strconv.Itoa(c[i]))
	}

	return res.String()
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
