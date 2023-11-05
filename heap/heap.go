package heap

import "golang.org/x/exp/constraints"

type Heap[T constraints.Ordered] struct {
	list      []T
	compareFn func(a, b T) bool
}

func NewMinHeap[T constraints.Ordered]() *Heap[T] {
	isSmaller := func(a, b T) bool {
		return a < b
	}

	h := &Heap[T]{
		compareFn: isSmaller,
	}

	return h
}

func NewMaxHeap[T constraints.Ordered]() *Heap[T] {
	isBigger := func(a, b T) bool {
		return a > b
	}

	h := &Heap[T]{
		compareFn: isBigger,
	}

	return h
}

func (h *Heap[T]) Top() T {
	return h.list[0]
}

func (h *Heap[T]) Push(val T) {
	h.list = append(h.list, val)
	h.up(len(h.list) - 1)
}

func (h *Heap[T]) Pop() {
	if len(h.list) <= 1 {
		h.list = nil
		return
	}
	h.swap(0, len(h.list)-1)
	h.list = h.list[:len(h.list)-1]
	h.down(0)
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.list) == 0
}

// Size returns the size of the heap
func (h *Heap[T]) Size() int {
	return len(h.list)
}

func (h *Heap[T]) swap(i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}

func (h *Heap[T]) up(child int) {
	if child <= 0 {
		return
	}
	parent := (child - 1) >> 1

	if !h.compareFn(h.list[child], h.list[parent]) {
		return
	}
	h.swap(child, parent)
	h.up(parent)
}

func (h *Heap[T]) down(parent int) {
	lessIdx := parent

	leftChild, rightChild := (parent*2)+1, (parent*2)+2
	if leftChild < len(h.list) && h.compareFn(h.list[leftChild], h.list[lessIdx]) {
		lessIdx = leftChild
	}
	if rightChild < len(h.list) && h.compareFn(h.list[rightChild], h.list[lessIdx]) {
		lessIdx = rightChild
	}
	if lessIdx == parent {
		return
	}

	h.swap(lessIdx, parent)
	h.down(lessIdx)
}
