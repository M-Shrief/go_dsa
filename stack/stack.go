package stack

import "github.com/M-Shrief/go-dsa-practice/linkedlist"

type Stack[T any] struct {
	list linkedlist.Singly[T]
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(val T) {
	stack.list.AddFirst(val)
	stack.size++
}

func (stack *Stack[T]) Pop() (T, bool) {
	val, ok := stack.list.DeleteFirst()
	if ok {
		stack.size--
	}
	return val, ok
}

func (stack *Stack[T]) GetTop() T {
	return stack.list.GetHead().GetVal()
}

func (stack *Stack[T]) GetSize() int {
	return stack.size
}
