package linkedlist

type DoublyNode[T any] struct {
	Val  T
	Next *DoublyNode[T]
	Prev *DoublyNode[T]
}

func NewDoublyNode[T any](val T) *DoublyNode[T] {
	return &DoublyNode[T]{val, nil, nil}
}
