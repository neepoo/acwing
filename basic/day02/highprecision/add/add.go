package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.acwing.com/problem/content/793/

func Add(a, b string) string {
	aLen := len(a)
	bLen := len(b)

	numA := make([]int, aLen)
	numB := make([]int, bLen)
	res := make([]int, 0)
	i := 0
	for j := aLen - 1; j >= 0; j-- {
		tmp, _ := strconv.Atoi(string(a[j]))
		numA[i] = tmp
		i++
	}
	i = 0
	for j := bLen - 1; j >= 0; j-- {
		tmp, _ := strconv.Atoi(string(b[j]))
		numB[i] = tmp
		i++
	}
	i, j := 0, 0
	c := 0 // 进喂
	for i < aLen && j < bLen {
		tmp := numA[i] + numB[j] + c
		res = append(res, tmp%10)
		c = tmp / 10
		i++
		j++
	}
	for i < aLen {
		tmp := numA[i] + c
		res = append(res, tmp%10)
		c = tmp / 10
		i++
	}
	for j < bLen {
		tmp := numB[j] + c
		res = append(res, tmp%10)
		c = tmp / 10
		j++
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
