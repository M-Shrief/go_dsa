package linkedlist

import (
	"fmt"
)

type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{val, nil}
}

func (n *Node[T]) GetVal() T {
	return n.val
}

func (n *Node[T]) GetNext() *Node[T] {
	return n.next
}

type Singly[T any] struct {
	Size int
	Head *Node[T]
}

func NewSingly[T any]() *Singly[T] {
	return &Singly[T]{}
}

func (list *Singly[T]) GetNode(pos int) any {

	if pos < 0 || pos >= list.Size {
		return nil
	}
	if pos == 0 {
		return list.Head.val
	}

	current := list.Head

	for count := 0; count < pos; count++ {
		current = current.next
	}

	return current.val
}

func (list *Singly[T]) AddFirst(val T) {
	n := NewNode(val)
	n.next = list.Head
	list.Head = n
	list.Size++
}

func (list *Singly[T]) AddLast(val T) {
	n := NewNode(val)

	if list.Head == nil {
		list.Head = n
		list.Size++
		return
	}

	current := list.Head
	for current.next != nil {
		current = current.next
	}
	current.next = n
	list.Size++
}

func (list *Singly[T]) DeleteFirst() (T, bool) {
	if list.Head == nil {
		var r T
		return r, false
	}

	current := list.Head
	list.Head = current.next
	list.Size--

	return current.val, true
}

func (list *Singly[T]) DeleteLast() (T, bool) {
	if list.Head == nil {
		var r T
		return r, false
	}

	if list.Head.next == nil {
		return list.DeleteFirst()
	}

	current := list.Head

	for current.next.next != nil {
		current = current.next
	}

	returnedValue := current.next.val
	current.next = nil
	list.Size--

	return returnedValue, true
}

func (list *Singly[T]) DeleteNode(pos int) (T, bool) {

	if pos < 0 || pos >= list.Size {
		var r T
		return r, false
	}
	if pos == 0 {
		return list.DeleteFirst()
	}
	if pos == list.Size-1 {
		return list.DeleteLast()
	}

	current := list.Head

	for count := 0; count < pos-1; count++ {
		current = current.next
	}

	removed := current.next.val
	current.next = current.next.next
	list.Size--

	return removed, true
}

func (list *Singly[T]) Reverse() {
	var prev, next *Node[T]
	current := list.Head

	for current != nil {
		next = current.next
		current.next = prev

		prev = current
		current = next
	}

	list.Head = prev
}

func (list *Singly[T]) Display() {
	for current := list.Head; current != nil; current = current.next {
		fmt.Println(current.val)
	}
}
