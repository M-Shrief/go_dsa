package tree

import (
	"fmt"

	"github.com/M-Shrief/go-dsa-practice/queue"
)

type BSNode[T any] struct {
	val   T
	left  *BSNode[T]
	right *BSNode[T]
}

func NewBSNode[T any](val T) *BSNode[T] {
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

type BST[T any] struct {
	root *BSNode[T]
	size int

	// Use standard [github.com/M-Shrief/go-dsa-practice/utils] for any type.field.
	// Usage:
	// 	compareFn := utis.Compare(a.age, b.age)
	// 	tree := NewBST[Person](compareFn)
	compareFn func(a, b T) int
}

func (bst *BST[T]) GetRoot() *BSNode[T] {
	return bst.root
}

func NewBST[T any](compareFn func(a, b T) int) *BST[T] {
	return &BST[T]{
		root:      nil,
		size:      0,
		compareFn: compareFn,
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
	if bst.compareFn(data, current.val) == -1 || bst.compareFn(data, current.val) == 0 {
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
	} else if bst.compareFn(data, node.val) == 0 {
		return node, true
	} else if bst.compareFn(data, node.val) == -1 {
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
	if bst.compareFn(data, node.val) == 0 {
		return nil
	} else if bst.compareFn(data, node.val) == -1 {
		if bst.compareFn(data, node.left.val) == 0 {
			return node
		} else {
			return bst.getNodeParent(node.left, data)
		}
	} else {
		if bst.compareFn(data, node.right.val) == 0 {
			return node
		} else {
			return bst.getNodeParent(node.right, data)
		}
	}
}

func (bst *BST[T]) Remove(data T) bool {
	if bst.size == 0 {
		return false
	} else if bst.size == 1 {
		bst.root = nil
		bst.size--
		return true
	} else {
		_, ok := bst.removeNode(bst.root, data)
		if ok {
			bst.size--
			return true
		} else {
			return false

		}
	}
}

func (bst *BST[T]) removeNode(node *BSNode[T], data T) (*BSNode[T], bool) {
	if node == nil {
		return node, false
	} else if bst.compareFn(data, node.val) == -1 { // if less, go left
		return bst.removeNode(node.left, data)
	} else if bst.compareFn(data, node.val) == 1 { // if more, go right
		return bst.removeNode(node.right, data)
	} else {
		ifLeftNodeExists := node.left != nil
		ifRightNodeExists := node.right != nil

		if ifLeftNodeExists && ifRightNodeExists {
			// // if it have 2 childs:
			// // set node.val to the maximum child in the left subtree.
			// // then go back and delete that leaf key.
			maximumValueInLeftSubTree, _ := bst.maximumNode(node.left)
			node.left, _ = bst.removeNode(node.left, maximumValueInLeftSubTree)
			node.val = maximumValueInLeftSubTree
			return node, true
		} else if ifLeftNodeExists {
			*node = *node.left
			return node, true
		} else if ifRightNodeExists {
			*node = *node.right
			return node, true
		} else {
			parent := bst.GetParent(node.val)

			if bst.compareFn(node.val, parent.val) == -1 || bst.compareFn(node.val, parent.val) == 0 {
				deletedNode := node.left
				parent.left = nil
				return deletedNode, true
			} else {
				deletedNode := node.right
				parent.right = nil
				return deletedNode, true
			}
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

func (bst *BST[T]) BFT() []T {

	if bst.GetSize() == 0 {
		return nil
	}

	list := make([]T, bst.GetSize())

	queue := queue.NewQueue[*BSNode[T]]()
	queue.Enqueue(bst.GetRoot())

	for queue.GetSize() > 0 {
		current, _ := queue.Dequeue()
		list = append(list, current.val)

		leftChild := current.GetLeft()
		if leftChild != nil {
			queue.Enqueue(leftChild)
		}

		rightChild := current.GetRight()
		if rightChild != nil {
			queue.Enqueue(rightChild)
		}
	}
	return list
}

func (bst *BST[T]) BFS(val T) (*BSNode[T], error) {
	var node *BSNode[T]

	if bst.GetSize() == 0 {
		return node, fmt.Errorf("BST is empty")
	}

	queue := queue.NewQueue[*BSNode[T]]()
	queue.Enqueue(bst.GetRoot())

	for queue.GetSize() > 0 {
		current, _ := queue.Dequeue()
		if bst.compareFn(current.val, val) == 0 {
			node = current
			break
		}

		leftChild := current.GetLeft()
		if leftChild != nil {
			queue.Enqueue(leftChild)
		}

		rightChild := current.GetRight()
		if rightChild != nil {
			queue.Enqueue(rightChild)
		}
	}
	return node, nil
}

type DFTMethod string

const (
	DFTPreOrder  DFTMethod = "preOrder"
	DFTInOrder   DFTMethod = "inOrder"
	DFTPostOrder DFTMethod = "postOrder"
)

/*
Depth First Traversal

takes DFTMethod as a parameter, equals a string from ("preOrder", "inOrder", "postOrder")

defaut traverse method: "inOrder"

you can import (DFTPreOrder, DFTInOrder, DFTPostOrder) as constant string.
*/
func (bst *BST[T]) DFT(method DFTMethod) []T {

	if bst.GetSize() == 0 {
		return nil
	}

	list := make([]T, bst.GetSize())

	switch method {
	case DFTPreOrder:
		return bst.preOrderTraversal(&list, bst.GetRoot())
	case DFTInOrder:
		return bst.inOrderTraversal(&list, bst.GetRoot())
	case DFTPostOrder:
		return bst.postOrderTraversal(&list, bst.GetRoot())
	default:
		return bst.inOrderTraversal(&list, bst.GetRoot())
	}
}

func (bst *BST[T]) preOrderTraversal(list *[]T, node *BSNode[T]) []T {
	if node == nil {
		return nil
	}

	*list = append(*list, node.GetVal())
	bst.preOrderTraversal(list, node.GetLeft())
	bst.preOrderTraversal(list, node.GetRight())

	return *list
}

func (bst *BST[T]) inOrderTraversal(list *[]T, node *BSNode[T]) []T {
	if node == nil {
		return nil
	}

	bst.inOrderTraversal(list, node.GetLeft())
	*list = append(*list, node.GetVal())
	bst.inOrderTraversal(list, node.GetRight())

	return *list
}

func (bst *BST[T]) postOrderTraversal(list *[]T, node *BSNode[T]) []T {
	if node == nil {
		return nil
	}

	bst.postOrderTraversal(list, node.GetLeft())
	bst.postOrderTraversal(list, node.GetRight())
	*list = append(*list, node.GetVal())

	return *list
}

/*
Depth First Search

takes DFTMethod as a parameter, equals a string from ("preOrder", "inOrder", "postOrder")

defaut traverse method: "inOrder"

you can import (DFTPreOrder, DFTInOrder, DFTPostOrder) as constant string.
*/
func (bst *BST[T]) DFS(data T, method DFTMethod) (*BSNode[T], error) {

	if bst.GetSize() == 0 {
		return nil, fmt.Errorf("tree is empty")
	}

	switch method {
	case DFTPreOrder:
		return bst.preOrderSearch(data, bst.GetRoot())
	case DFTInOrder:
		return bst.inOrderSearch(data, bst.GetRoot())
	case DFTPostOrder:
		return bst.postOrderSearch(data, bst.GetRoot())
	default:
		return bst.inOrderSearch(data, bst.GetRoot())
	}
}

func (bst *BST[T]) preOrderSearch(data T, node *BSNode[T]) (*BSNode[T], error) {
	if node == nil {
		return nil, fmt.Errorf("nil node")
	}
	if bst.compareFn(data, node.val) == 0 {
		return node, nil
	}

	n1, err := bst.preOrderSearch(data, node.GetLeft())
	if err == nil {
		return n1, nil
	}

	n2, err := bst.preOrderSearch(data, node.GetRight())
	if err == nil {
		return n2, nil
	}

	return nil, fmt.Errorf("%v doesn't exist", data)
}

func (bst *BST[T]) inOrderSearch(data T, node *BSNode[T]) (*BSNode[T], error) {
	if node == nil {
		return nil, fmt.Errorf("nil node")
	}

	n1, err := bst.inOrderSearch(data, node.GetLeft())
	if err == nil {
		return n1, nil
	}

	if bst.compareFn(data, node.val) == 0 {
		return node, nil
	}

	n2, err := bst.inOrderSearch(data, node.GetRight())
	if err == nil {
		return n2, nil
	}

	return nil, fmt.Errorf("%v doesn't exist", data)
}

func (bst *BST[T]) postOrderSearch(data T, node *BSNode[T]) (*BSNode[T], error) {
	if node == nil {
		return nil, fmt.Errorf("nil node")
	}

	n1, err := bst.postOrderSearch(data, node.GetLeft())
	if err == nil {
		return n1, nil
	}

	n2, err := bst.postOrderSearch(data, node.GetRight())
	if err == nil {
		return n2, nil
	}

	if bst.compareFn(data, node.val) == 0 {
		return node, nil
	}

	return nil, fmt.Errorf("%v doesn't exist", data)
}
