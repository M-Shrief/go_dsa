package heap

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

type Heap[T any] struct {
	list      []T
	compareFn func(a, b T) bool
}

type CResult int

const (
	CR1 CResult = 1
	CR2 CResult = -1
	CR3 CResult = 0
)

func Compare[T constraints.Ordered](a, b T) CResult {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func NewHeap[T any](compareFn func(a, b T) bool) *Heap[T] {
	return &Heap[T]{compareFn: compareFn}
}

func (h *Heap[T]) Top() T {
	return h.list[0]
}

func (h *Heap[T]) Push(val T) {
	h.list = append(h.list, val)
	h.up(len(h.list) - 1)
}

// Delete by value
func (h *Heap[T]) Delete(val T) {
	i := 0
	for !reflect.DeepEqual(h.list[i], val) {
		i++
	}
	h.list = append(h.list[:i], h.list[i+1:]...)
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
