package main
// https://www.acwing.com/problem/content/150/
import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N = 100010

type minHeap []int

var (
	res, n int
	mh     minHeap
)

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	old := *m
	n := m.Len()
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)
	var x int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x)
		mh = append(mh, x)
	}
	heap.Init(&mh)
	for n > 1 {  // 进行n-1次操作
		m1, m2 := heap.Pop(&mh).(int), heap.Pop(&mh).(int)
		res += m1 + m2
		heap.Push(&mh, m1+m2)
		n--
	}
	fmt.Println(res)
}
