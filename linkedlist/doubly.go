package linkedlist

import "fmt"

type DoublyNode[T any] struct {
	val  T
	next *DoublyNode[T]
	prev *DoublyNode[T]
}

func NewDoublyNode[T any](val T) *DoublyNode[T] {
	return &DoublyNode[T]{val, nil, nil}
}

func (n *DoublyNode[T]) GetVal() T {
	return n.val
}

func (n *DoublyNode[T]) GetNext() *DoublyNode[T] {
	return n.next
}

func (n *DoublyNode[T]) GetPrev() *DoublyNode[T] {
	return n.prev
}

type DoublyLinkedList[T any] struct {
	Size int
	Head *DoublyNode[T]
	Tail *DoublyNode[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (list *DoublyLinkedList[T]) GetHead() *DoublyNode[T] {
	return list.Head
}

func (list *DoublyLinkedList[T]) GetTail() *DoublyNode[T] {
	if list.Tail == nil {
		return list.Head
	}
	return list.Tail
}

func (list *DoublyLinkedList[T]) GetNode(pos int) any {
	if pos < 0 || pos >= list.Size {
		return nil
	}
	if pos == 0 {
		return list.Head
	}
	if pos == list.Size-1 {
		return list.Tail
	}

	current := list.Head

	for count := 0; count < pos; count++ {
		current = current.next
	}

	return current
}

func (list *DoublyLinkedList[T]) AddFirst(val T) {
	n := NewDoublyNode(val)
	if list.Head == nil {
		list.Head = n
		list.Size++
		return
	}
	n.next = list.Head
	list.Head.prev = n
	list.Head = n
	if list.Size == 1 {
		list.Tail = n.next
	}
	list.Size++
}

func (list *DoublyLinkedList[T]) AddLast(val T) {
	if list.Size == 0 {
		list.AddFirst(val)
		return
	}
	n := NewDoublyNode(val)
	list.GetTail().next = n
	n.prev = list.GetTail()
	list.Tail = n
	list.Size++
}

func (list *DoublyLinkedList[T]) DeleteFirst() (T, bool) {
	if list.Size == 0 {
		var r T
		return r, false
	}

	node := list.Head

	if list.Size == 1 {
		list.Head = nil
		list.Tail = nil
		list.Size--
		return node.val, true
	}

	list.Head = node.next
	list.Head.prev = nil
	list.Size--

	return node.val, true
}

func (list *DoublyLinkedList[T]) DeleteLast() (T, bool) {
	if list.Size == 0 {
		var r T
		return r, false
	}

	if list.Size == 1 {
		return list.DeleteFirst()
	}

	node := list.Tail
	list.Tail = node.prev
	list.Tail.next = nil
	list.Size--

	return node.val, true
}

func (list *DoublyLinkedList[T]) DeleteNode(pos int) (T, bool) {
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

	removedVal := current.next.val
	current.next = current.next.next
	current.next.prev = current
	list.Size--

	return removedVal, true

}

func (list *DoublyLinkedList[T]) Display() {
	for current := list.Head; current != nil; current = current.next {
		fmt.Println(current.val)
	}
}
