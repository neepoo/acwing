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
	s, b := Div(a, b)
	fmt.Printf("%s\n%d", s, b)
}

func Div(a string, b int) (string, int) {
	return div(sToSlice(a), b)
}

func reverse(a []int) []int {
	i, j := 0, len(a)-1

	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	return a
}

func div(a []int, b int) (string, int) {
	t := 0
	var c []int
	res := strings.Builder{}
	for i := len(a) - 1; i >= 0; i-- {
		t = t*10 + a[i]
		c = append(c, t/b)
		t %= b
	}
	c = reverse(c)
	// 去除前导0
	for i := len(c) - 1; i >= 1 && c[i] == 0; i-- {
		c = c[:len(c)-1]
	}
	for i := len(c) - 1; i >= 0; i-- {
		res.WriteString(strconv.Itoa(c[i]))
	}

	return res.String(), t
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
