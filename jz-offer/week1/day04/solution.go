package main

import "fmt"

func replaceSpaces(s string) string {
	var res string
	i, j := 0, 0
	for i < len(s) {
		for i < len(s) && s[i] != ' ' {
			i++
		}
		// 现在i是第一个空格开始的地方
		res += s[j:i]
		j = i
		for j < len(s) && s[j] == ' ' {
			j++
		}
		// j现在是第一个不是空格的地方
		for k := i; k < j; k++ {
			res += "%20"
		}
		i = j
	}
	return res
}

func main()  {
	s := "We are happy.  "
	s1 := replaceSpaces(s)
	fmt.Println(s)
	fmt.Println(s1)
}