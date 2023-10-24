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

func (bst *BST[T]) Insert(data T) {
	if bst.size == 0 {
		newNode := NewBSNode[T](data)
		bst.root = newNode
		bst.size++
	} else {
		bst.insertNode(bst.root, data)
	}
}

func (bst *BST[T]) insertNode(current *BSNode[T], data T) {
	if data <= current.val {
		if current.left != nil {
			bst.insertNode(current.left, data)
		} else {
			newNode := NewBSNode[T](data)
			current.left = newNode
			bst.size++
		}
	} else {
		if current.right != nil {
			bst.insertNode(current.right, data)
		} else {
			newNode := NewBSNode[T](data)
			current.right = newNode
			bst.size++
		}
	}
}

func (bst *BST[T]) Search(data T) (*BSNode[T], bool) {
	if bst.size == 0 {
		var r T
		newNode := NewBSNode[T](r)
		return newNode, false
	} else {
		return bst.searchNode(bst.root, data)
	}
}

func (bst *BST[T]) searchNode(node *BSNode[T], data T) (*BSNode[T], bool) {
	if node == nil {
		var r T
		newNode := NewBSNode[T](r)
		return newNode, false
	} else if node.GetVal() == data {
		return node, true
	} else if data < node.GetVal() {
		return bst.searchNode(node.left, data)
	} else {
		return bst.searchNode(node.right, data)
	}
}
