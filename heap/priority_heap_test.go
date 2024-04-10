package heap

import "testing"

func TestPriorityHeap(t *testing.T) {
	t.Run("Testing Push()", func(t *testing.T) {
		heap := NewPriorityHeap[int]()

		heap.Push(8, 7)
		heap.Push(6, 11)
		heap.Push(4, 12)

		top1 := heap.Peak()
		if top1 != 8 {
			t.Errorf("Got: %v, Want: %v", top1, 8)
		}

		heap.Push(20, 6)
		heap.Push(24, 16)
		heap.Push(23, 18)
		heap.Push(21, 19)

		top2 := heap.Peak()
		if top2 != 20 {
			t.Errorf("Got: %v, Want: %v", top2, 20)
		}

		if !(heap.list[1].priority < heap.list[(1*2)+1].priority && heap.list[1].priority < heap.list[(1*2)+2].priority) {
			t.Errorf("Index: %v is not smaller than it's childs", heap.list[1])
		}
		if !(heap.list[2].priority < heap.list[(2*2)+1].priority && heap.list[2].priority < heap.list[(2*2)+2].priority) {
			t.Errorf("Index: %v is not smaller than it's childs", heap.list[2])
		}
	})

	t.Run("Testing Pop()", func(t *testing.T) {
		heap := NewPriorityHeap[int]()
		heap.Push(8, 7)
		heap.Push(6, 11)
		heap.Push(4, 12)

		pop1 := heap.Pop()

		if pop1 != 8 {
			t.Errorf("Got: %v, Want: %v", pop1, 8)
		}

		top1 := heap.Peak()
		if top1 == 8 && top1 != 6 {
			t.Errorf("Got: %v, Want: %v", top1, 6)
		}

		heap.Push(20, 6)
		heap.Push(24, 16)
		heap.Push(23, 18)
		heap.Push(21, 19)

		pop2 := heap.Pop()

		if pop2 != 20 {
			t.Errorf("Got: %v, Want: %v", pop2, 20)
		}

		top2 := heap.Peak()
		if top2 == 20 && top2 != 6 {
			t.Errorf("Got: %v, Want: %v", top2, 6)
		}

		heap.Push(200, 4)
		heap.Push(12, 3)

		if !(heap.list[1].priority < heap.list[(1*2)+1].priority && heap.list[1].priority < heap.list[(1*2)+2].priority) {
			t.Errorf("Index: %v is not smaller than it's childs", heap.list[1])
		}
		if !(heap.list[2].priority < heap.list[(2*2)+1].priority && heap.list[2].priority < heap.list[(2*2)+2].priority) {
			t.Errorf("Index: %v is not smaller than it's childs", heap.list[2])
		}
	})
}
