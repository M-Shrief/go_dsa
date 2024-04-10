package heap

type node[T any] struct {
	value T
	// the smaller the more it's prioritized
	priority uint64
}

func newNode[T any](value T, priority uint64) *node[T] {
	return &node[T]{value, priority}
}

type PriorityHeap[T any] struct {
	list []*node[T]
}

func (ph *PriorityHeap[T]) isPreceding(a, b uint64) bool {
	return a < b
}

func NewPriorityHeap[T any]() *PriorityHeap[T] {
	return &PriorityHeap[T]{}
}

func (h *PriorityHeap[T]) Peak() T {
	return h.list[0].value
}

/*
Add new item for the Heap.

the smaller the priority value, the more it's prioritized
*/
func (h *PriorityHeap[T]) Push(value T, priority uint64) {
	n := newNode[T](value, priority)
	h.list = append(h.list, n)
	h.up(len(h.list) - 1)
}

func (h *PriorityHeap[T]) Pop() T {
	if len(h.list) <= 1 {
		h.list = nil
		var r T
		return r
	}
	top := h.Peak()

	h.swap(0, len(h.list)-1)
	h.list = h.list[:len(h.list)-1]
	h.down(0)

	return top
}

func (h *PriorityHeap[T]) swap(i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}

func (h *PriorityHeap[T]) up(child int) {
	if child <= 0 {
		return
	}
	parent := (child - 1) >> 1

	if !h.isPreceding(h.list[child].priority, h.list[parent].priority) {
		return
	}
	h.swap(child, parent)
	h.up(parent)
}

func (h *PriorityHeap[T]) down(parent int) {
	lessIdx := parent

	leftChild, rightChild := (parent*2)+1, (parent*2)+2
	if leftChild < len(h.list) && h.isPreceding(h.list[leftChild].priority, h.list[lessIdx].priority) {
		lessIdx = leftChild
	}
	if rightChild < len(h.list) && h.isPreceding(h.list[rightChild].priority, h.list[lessIdx].priority) {
		lessIdx = rightChild
	}
	if lessIdx == parent {
		return
	}

	h.swap(lessIdx, parent)
	h.down(lessIdx)
}
