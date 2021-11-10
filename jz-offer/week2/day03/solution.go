package day03

func NumberOf1(n int) int {
	var res int
	for i:=0;i<32;i++{
		res+=(n >> i) & 1
	}
	return res
}

//func NumberOf1(n int) int {
//	var count uint
//	n1 := uint(n)
//	for n1 > 0{
//		count+= n1 & 1
//		n1 >>=1
//	}
//	return int(count)
//}
//
//func NumberOf1(n int) int {
//	count:=0
//	for n!=0 {
//		n = n & n-1
//		count++
//	}
//	return count
//}