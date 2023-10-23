package linkedlist

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{val, nil}
}

type DoublyNode[T any] struct {
	Val  T
	Next *DoublyNode[T]
	Prev *DoublyNode[T]
}

func NewDoublyNode[T any](val T) *DoublyNode[T] {
	return &DoublyNode[T]{val, nil, nil}
}
