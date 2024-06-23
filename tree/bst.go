package tree

import (
	"fmt"

	"github.com/M-Shrief/go-dsa-practice/queue"
	"golang.org/x/exp/constraints"
)

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
	} else if data < node.val { // if less, go left
		return bst.removeNode(node.left, data)
	} else if data > node.val { // if more, go right
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

			if node.val <= parent.val {
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
		if current.val == val {
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
		return preOrderTraversal(&list, bst.GetRoot())
	case DFTInOrder:
		return inOrderTraversal(&list, bst.GetRoot())
	case DFTPostOrder:
		return postOrderTraversal(&list, bst.GetRoot())
	default:
		return inOrderTraversal(&list, bst.GetRoot())
	}
}

func preOrderTraversal[T constraints.Ordered](list *[]T, node *BSNode[T]) []T {
	if node == nil {
		return nil
	}

	*list = append(*list, node.GetVal())
	preOrderTraversal[T](list, node.GetLeft())
	preOrderTraversal[T](list, node.GetRight())

	return *list
}

func inOrderTraversal[T constraints.Ordered](list *[]T, node *BSNode[T]) []T {
	if node == nil {
		return nil
	}

	inOrderTraversal[T](list, node.GetLeft())
	*list = append(*list, node.GetVal())
	inOrderTraversal[T](list, node.GetRight())

	return *list
}

func postOrderTraversal[T constraints.Ordered](list *[]T, node *BSNode[T]) []T {
	if node == nil {
		return nil
	}

	postOrderTraversal[T](list, node.GetLeft())
	postOrderTraversal[T](list, node.GetRight())
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
		return preOrderSearch(data, bst.GetRoot())
	case DFTInOrder:
		return inOrderSearch(data, bst.GetRoot())
	case DFTPostOrder:
		return postOrderSearch(data, bst.GetRoot())
	default:
		return inOrderSearch(data, bst.GetRoot())
	}
}

func preOrderSearch[T constraints.Ordered](data T, node *BSNode[T]) (*BSNode[T], error) {
	if node == nil {
		return nil, fmt.Errorf("nil node")
	}
	if data == node.val {
		return node, nil
	}

	n1, err := preOrderSearch[T](data, node.GetLeft())
	if err == nil {
		return n1, nil
	}

	n2, err := preOrderSearch[T](data, node.GetRight())
	if err == nil {
		return n2, nil
	}

	return nil, fmt.Errorf("%v doesn't exist", data)
}

func inOrderSearch[T constraints.Ordered](data T, node *BSNode[T]) (*BSNode[T], error) {
	if node == nil {
		return nil, fmt.Errorf("nil node")
	}

	n1, err := inOrderSearch[T](data, node.GetLeft())
	if err == nil {
		return n1, nil
	}

	if data == node.val {
		return node, nil
	}

	n2, err := inOrderSearch[T](data, node.GetRight())
	if err == nil {
		return n2, nil
	}

	return nil, fmt.Errorf("%v doesn't exist", data)
}

func postOrderSearch[T constraints.Ordered](data T, node *BSNode[T]) (*BSNode[T], error) {
	if node == nil {
		return nil, fmt.Errorf("nil node")
	}

	n1, err := postOrderSearch[T](data, node.GetLeft())
	if err == nil {
		return n1, nil
	}

	n2, err := postOrderSearch[T](data, node.GetRight())
	if err == nil {
		return n2, nil
	}

	if data == node.val {
		return node, nil
	}

	return nil, fmt.Errorf("%v doesn't exist", data)
}
