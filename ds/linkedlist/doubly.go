package linkedlist

import "fmt"

// Node for Doubly LinkedList
type DoublyNode[T any] struct {
	// Node Value
	val T
	// Next node in LinkedList
	next *DoublyNode[T]
	// Previous node in LinkedList
	prev *DoublyNode[T]
}

// Create New DoublyNode, and assign its value
func NewDoublyNode[T any](val T) *DoublyNode[T] {
	return &DoublyNode[T]{val, nil, nil}
}

// Get Node's value
func (n *DoublyNode[T]) GetVal() T {
	return n.val
}

// Get the next node
func (n *DoublyNode[T]) GetNext() *DoublyNode[T] {
	return n.next
}

// Get the previous node
func (n *DoublyNode[T]) GetPrev() *DoublyNode[T] {
	return n.prev
}

// Doubly Linkedist
type Doubly[T any] struct {
	size int
	head *DoublyNode[T]
	tail *DoublyNode[T]
}

// Create A new Doubly LinkedList
func NewDoubly[T any]() *Doubly[T] {
	return &Doubly[T]{}
}

// Get LinkedList's Size
func (list *Doubly[T]) GetSize() int {
	return list.size
}

// Get LinkedList's Head
func (list *Doubly[T]) GetHead() *DoublyNode[T] {
	return list.head
}

// Get LinkedList's Tail
func (list *Doubly[T]) GetTail() *DoublyNode[T] {
	if list.tail == nil {
		return list.head
	}
	return list.tail
}

/*
Get specific node by position (pos)
*/
func (list *Doubly[T]) GetNode(pos int) *DoublyNode[T] {
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

// Add value as a NewDoublyNode to Doubly LinkedList's First position.
func (list *Doubly[T]) AddFirst(val T) {
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

// Add value as a NewDoublyNode to Doubly LinkedList's Last position.
func (list *Doubly[T]) AddLast(val T) {
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

/*
Delete First node from LinkedList.

Returns deletedNode's value and boolean to indicate if delete is successful or not.
*/
func (list *Doubly[T]) DeleteFirst() (T, bool) {
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

/*
Delete Last node from LinkedList.

Returns deletedNode's value and boolean to indicate if delete is successful or not.
*/
func (list *Doubly[T]) DeleteLast() (T, bool) {
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

/*
Delete node by position from LinkedList.

Returns deletedNode's value and boolean to indicate if delete is successful or not.
*/
func (list *Doubly[T]) DeleteByPosition(pos int) (T, bool) {
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

/*
Delete by prividing the target DoublyNode from LinkedList

Returns deletedNode's value.
*/
func (list *Doubly[T]) DeleteByNode(node *DoublyNode[T]) T {
	if node == list.head {
		val, _ := list.DeleteFirst()
		return val
	}
	if node == list.tail {
		val, _ := list.DeleteLast()
		return val
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil

	list.size--
	return node.val
}

// Print LinkedList nodes in order.
func (list *Doubly[T]) Display() {
	for current := list.head; current != nil; current = current.next {
		fmt.Println(current.val)
	}
}
