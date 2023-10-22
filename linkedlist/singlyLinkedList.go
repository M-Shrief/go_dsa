package linkedlist

import (
	"fmt"
)

type SinglyLinkedList[T any] struct {
	Size int
	Head *Node[T]
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (list *SinglyLinkedList[T]) GetNode(pos int) any {

	if pos < 0 || pos >= list.Size {
		return nil
	}
	if pos == 0 {
		return list.Head.Val
	}

	current := list.Head

	for count := 0; count < pos; count++ {
		current = current.Next
	}

	return current.Val
}

func (list *SinglyLinkedList[T]) AddFirst(val T) {
	n := NewNode(val)
	n.Next = list.Head
	list.Head = n
	list.Size++
}

func (list *SinglyLinkedList[T]) AddLast(val T) {
	n := NewNode(val)

	if list.Head == nil {
		list.Head = n
		list.Size++
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = n
	list.Size++
}

func (list *SinglyLinkedList[T]) DeleteFirst() (T, bool) {
	if list.Head == nil {
		var r T
		return r, false
	}

	current := list.Head
	list.Head = current.Next
	list.Size--

	return current.Val, true
}

func (list *SinglyLinkedList[T]) DeleteLast() (T, bool) {
	if list.Head == nil {
		var r T
		return r, false
	}

	if list.Head.Next == nil {
		return list.DeleteFirst()
	}

	current := list.Head

	for current.Next.Next != nil {
		current = current.Next
	}

	returnedValue := current.Next.Val
	current.Next = nil
	list.Size--

	return returnedValue, true
}

func (list *SinglyLinkedList[T]) DeleteNode(pos int) (T, bool) {

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
		current = current.Next
	}

	removed := current.Next.Val
	current.Next.Next = current.Next
	list.Size--

	return removed, true
}

func (list *SinglyLinkedList[T]) Reverse() {
	var prev, next *Node[T]
	current := list.Head

	for current != nil {
		next = current.Next
		current.Next = prev

		prev = current
		current = next
	}

	list.Head = prev
}

func (list *SinglyLinkedList[T]) Display() {
	for current := list.Head; current != nil; current = current.Next {
		fmt.Println(current.Val)
	}
}
