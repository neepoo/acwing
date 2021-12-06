package main

import (
	"fmt"
	"strings"
)

//func reverseWords(s string) string {
//	var res string
//	n := len(s)
//	j := n - 1
//	for i := n - 1; i >= 0; {
//		for i >= 0 && s[i] == ' ' {
//			i--
//		}
//		j = i
//		for j >= 0 && s[j] != ' ' {
//			j--
//		}
//		if len(res) > 0 {
//			res += " " + s[j+1:i+1]
//		} else {
//			res += s[j+1 : i+1]
//		}
//
//		i = j
//		for i >= 0 && s[i] == ' ' {
//			i--
//		}
//		j = i - 1
//	}
//
//	return res
//}

func reverseWords(s string) string {
	ss := strings.Fields(s)
	l, r := 0, len(ss)-1
	for l < r {
		ss[l], ss[r] = ss[r], ss[l]
		l++
		r--
	}
	return strings.Join(ss, " ")

}

// 58 -1
func reverseLeftWords(s string, k int) string {
	n := len(s)
	if k > n {
		return ""
	}
	t1 := s[:k]
	t2 := s[k:]
	return t2 + t1
}
func main() {
	fmt.Println(reverseWords("   the sky is blue   "))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("a good   example"))

	fmt.Println(reverseLeftWords("abcdefg", 2))
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
	/*
		n += wasSpace & ^isSpace
	*/
	isSpace := 65535
	a := ^isSpace
	fmt.Printf("isSpcae= %32b, a=%32b\n", isSpace, a)
	//fmt.Println(strconv.FormatInt(int64(isSpace), 2), strconv.FormatInt(int64(a), 2))
	fmt.Println(a)      // -1
	fmt.Println(1 & ^0) // 1
	fmt.Println(1 & -2)
	fmt.Println(1 & -1) // 1
}
