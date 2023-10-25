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
		return nil, false
	} else {
		return bst.searchNode(bst.root, data)
	}
}

func (bst *BST[T]) searchNode(node *BSNode[T], data T) (*BSNode[T], bool) {
	if node == nil {
		return nil, false
	} else if node.val == data {
		return node, true
	} else if data < node.val {
		return bst.searchNode(node.left, data)
	} else {
		return bst.searchNode(node.right, data)
	}
}

func (bst *BST[T]) GetParent(data T) *BSNode[T] {
	if bst.size == 0 {
		return nil
	} else {
		return bst.getNodeParent(bst.root, data)
	}
}

func (bst *BST[T]) getNodeParent(node *BSNode[T], data T) *BSNode[T] {
	if data == node.val {
		return nil
	} else if data < node.val {
		if data == node.left.val {
			return node
		} else {
			return bst.getNodeParent(node.left, data)
		}
	} else {
		if data == node.right.val {
			return node
		} else {
			return bst.getNodeParent(node.right, data)
		}
	}
}

func (bst *BST[T]) Minimum() (T, bool) {
	if bst.size == 0 {
		var r T
		return r, false
	}
	return bst.minimumNode(bst.root)
}

func (bst *BST[T]) minimumNode(node *BSNode[T]) (T, bool) {
	if node == nil {
		var r T
		return r, false
	} else if node.left == nil {
		return node.val, true
	} else {
		return bst.minimumNode(node.left)
	}
}

func (bst *BST[T]) Maximum() (T, bool) {
	if bst.size == 0 {
		var r T
		return r, false
	}
	return bst.maximumNode(bst.root)
}

func (bst *BST[T]) maximumNode(node *BSNode[T]) (T, bool) {
	if node == nil {
		var r T
		return r, false
	} else if node.right == nil {
		return node.val, true
	} else {
		return bst.maximumNode(node.right)
	}
}
