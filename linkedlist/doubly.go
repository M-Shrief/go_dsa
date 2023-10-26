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
	size int
	head *DoublyNode[T]
	tail *DoublyNode[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (list *DoublyLinkedList[T]) GetSize() int {
	return list.size
}

func (list *DoublyLinkedList[T]) GetHead() *DoublyNode[T] {
	return list.head
}

func (list *DoublyLinkedList[T]) GetTail() *DoublyNode[T] {
	if list.tail == nil {
		return list.head
	}
	return list.tail
}

func (list *DoublyLinkedList[T]) GetNode(pos int) *DoublyNode[T] {
	if pos < 0 || pos >= list.size {
		return nil
	}
	if pos == 0 {
		return list.head
	}
	if pos == list.size-1 {
		return list.tail
	}

	current := list.head

	for count := 0; count < pos; count++ {
		current = current.next
	}

	return current
}

func (list *DoublyLinkedList[T]) AddFirst(val T) {
	n := NewDoublyNode(val)
	if list.head == nil {
		list.head = n
		list.size++
		return
	}
	n.next = list.head
	list.head.prev = n
	list.head = n
	if list.size == 1 {
		list.tail = n.next
	}
	list.size++
}

func (list *DoublyLinkedList[T]) AddLast(val T) {
	if list.size == 0 {
		list.AddFirst(val)
		return
	}
	n := NewDoublyNode(val)
	list.GetTail().next = n
	n.prev = list.GetTail()
	list.tail = n
	list.size++
}

func (list *DoublyLinkedList[T]) DeleteFirst() (T, bool) {
	if list.size == 0 {
		var r T
		return r, false
	}

	node := list.head

	if list.size == 1 {
		list.head = nil
		list.tail = nil
		list.size--
		return node.val, true
	}

	list.head = node.next
	list.head.prev = nil
	list.size--

	return node.val, true
}

func (list *DoublyLinkedList[T]) DeleteLast() (T, bool) {
	if list.size == 0 {
		var r T
		return r, false
	}

	if list.size == 1 {
		return list.DeleteFirst()
	}

	node := list.tail
	list.tail = node.prev
	list.tail.next = nil
	list.size--

	return node.val, true
}

func (list *DoublyLinkedList[T]) DeleteByPosition(pos int) (T, bool) {
	if pos < 0 || pos >= list.size {
		var r T
		return r, false
	}
	if pos == 0 {
		return list.DeleteFirst()
	}
	if pos == list.size-1 {
		return list.DeleteLast()
	}

	current := list.head
	for count := 0; count < pos-1; count++ {
		current = current.next
	}

	removedVal := current.next.val
	current.next = current.next.next
	current.next.prev = current
	list.size--

	return removedVal, true

}

func (list *DoublyLinkedList[T]) Display() {
	for current := list.head; current != nil; current = current.next {
		fmt.Println(current.val)
	}
}
