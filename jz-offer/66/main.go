package main

func constructArr(a []int) []int {
	n := len(a)
	b := make([]int, n)
	p := 1
	for i := 0; i < n; i++ {
		b[i] = p
		p *= a[i]
	}
	p = 1
	for i := n - 1; i >= 0; i-- {
		b[i] *= p
		p *= a[i]
	}
	return b
}

func main() {

}
