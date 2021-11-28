package main

import "fmt"

type Stack struct {
	ns [20000]int
	tt int
}

func NewStack() *Stack {
	return &Stack{}
}
func (s *Stack) pop() {
	s.tt--
}

func (s *Stack) push(x int) {
	s.tt++
	s.ns[s.tt] = x

}

func (s *Stack) top() int {
	return s.ns[s.tt]
}

func (s *Stack) empty() bool {
	return s.tt == 0
}

type MinStack struct {
	ms *Stack
	vs *Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		ms: NewStack(),
		vs: NewStack(),
	}
}

func (s *MinStack) Push(x int) {
	s.vs.push(x)
	if s.ms.empty() || s.ms.top() >= x{
		s.ms.push(x)
	}
}

func (s *MinStack) Pop() {
	if s.ms.top() == s.vs.top(){
		s.ms.pop()
	}
	s.vs.pop()
}

func (s *MinStack) Top() int {
	return s.vs.top()
}

func (s *MinStack) Min() int {
	return s.ms.top()
}

func main() {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.Min())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.Min())
}
