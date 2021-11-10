package main

import (
	"fmt"
	"strconv"
)

func isNumber(s string) bool {
	i, err :=strconv.Atoi(s)
	fmt.Println(i, err)
	return false
}

func main()  {
	isNumber("+22")
	isNumber("-22")
}
