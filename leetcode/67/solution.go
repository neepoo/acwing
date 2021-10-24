package solution

import (
	"strconv"
)

func addBinary(a string, b string) string {
	var reverseRes string
	c := 0
	var i int
	for i = 0; i < len(a) && i < len(b); i++ {
		curA := a[len(a)-i-1]
		curB := b[len(b)-i-1]
		curAInt := 0
		if curA == '1' {
			curAInt = 1
		}
		curBInt := 0
		if curB == '1' {
			curBInt = 1
		}
		curAns := curAInt + curBInt + c
		c = curAns / 2
		reverseRes += strconv.Itoa(curAns % 2)

	}
	for j := i; j < len(a); j++ {
		curA := a[len(a)-j-1]
		curAInt := 0
		if curA == '1' {
			curAInt = 1
		}
		curAns := curAInt + c
		c = curAns / 2
		reverseRes += strconv.Itoa(curAns % 2)
	}
	for j := i; j < len(b); j++ {
		curB := b[len(b)-j-1]
		curBInt := 0
		if curB == '1' {
			curBInt = 1
		}
		curAns := curBInt + c
		c = curAns / 2
		reverseRes += strconv.Itoa(curAns % 2)
	}
	if c >= 1 {
		reverseRes += "1"
	}
	res := ""
	for k := len(reverseRes) - 1; k >= 0; k-- {
		res += string(reverseRes[k])
	}
	return res
}

func addBinaryNew(a string, b string) string {
	var reverseRes string
	c := 0
	var i int
	for i = 0; i < len(a) || i < len(b); i++ {
		curAInt := 0
		curBInt := 0
		if len(a)-i-1 >= 0 {
			curA := a[len(a)-i-1]
			if curA == '1' {
				curAInt = 1
			}

		}
		if len(b)-i-1 >= 0 {
			curB := b[len(b)-i-1]
			if curB == '1' {
				curBInt = 1
			}
		}
		curAns := curAInt + curBInt + c
		c = curAns / 2
		reverseRes += strconv.Itoa(curAns % 2)

	}
	if c >= 1 {
		reverseRes += "1"
	}
	res := ""
	for k := len(reverseRes) - 1; k >= 0; k-- {
		res += string(reverseRes[k])
	}
	return res
}
