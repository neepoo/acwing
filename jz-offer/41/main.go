package main

import "container/heap"

type MinHeap []int
type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {

	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	*m = old[:n-1]
	return old[n-1]
}
func (m *MaxHeap) top() int {
	ans := heap.Pop(m).(int)
	heap.Push(m, ans)
	return ans
}

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	*m = old[:n-1]
	return old[n-1]
}

func (m *MinHeap) top() int {
	ans := heap.Pop(m).(int)
	heap.Push(m, ans)
	return ans
}

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxHeap, num)
	for this.minHeap.Len() > 0 && this.minHeap.top() < this.maxHeap.top() {
		minv, maxv := heap.Pop(this.minHeap).(int), heap.Pop(this.maxHeap).(int)
		heap.Push(this.maxHeap, minv)
		heap.Push(this.minHeap, maxv)
	}
	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.maxHeap.Len()+this.minHeap.Len())%2 == 1 {
		return float64(this.maxHeap.top())
	}
	return (float64(this.maxHeap.top()) + float64(this.minHeap.top())) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {

}
