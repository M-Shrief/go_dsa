package tree

import "golang.org/x/exp/constraints"

type BSNode[T constraints.Ordered] struct {
	val   T
	left  *BSNode[T]
	right *BSNode[T]
}

func NewBSNode[T constraints.Ordered](val T) *BSNode[T] {
	return &BSNode[T]{val, nil, nil}
}

func (n *BSNode[T]) GetVal() T {
	return n.val
}

func (n *BSNode[T]) GetLeft() *BSNode[T] {
	return n.left
}

func (n *BSNode[T]) GetRight() *BSNode[T] {
	return n.right
}

type BST[T constraints.Ordered] struct {
	root *BSNode[T]
	size int
}

func (bst *BST[T]) GetRoot() *BSNode[T] {
	return bst.root
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{
		root: nil,
		size: 0,
	}
}

func (bst *BST[T]) GetSize() int {
	return bst.size
}
