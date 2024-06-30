package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	t.Run("MinHeap: Testing Push()", func(t *testing.T) {
		isSmaller := func(a, b int) bool {
			if Compare(a, b) == CR2 {
				return true
			} else {
				return false
			}
		}
		heap := &Heap[int]{[]int{}, isSmaller}
		heap.Push(8)
		heap.Push(6)
		heap.Push(4)
		if heap.Top() != 4 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 4)
		}
		heap.Push(3)
		heap.Push(0)
		heap.Push(-2)
		heap.Push(-5)
		if heap.Top() != -5 {
			t.Errorf("Got: %v, Want: %v", heap.list, -5)
		}
		if !(heap.list[1] < heap.list[(1*2)+1] && heap.list[1] < heap.list[(1*2)+2]) {
			t.Errorf("Index: %v is not smaller than it's childs", heap.list[1])
		}
		if !(heap.list[2] < heap.list[(2*2)+1] && heap.list[2] < heap.list[(2*2)+2]) {
			t.Errorf("Index: %v is not smaller than it's childs", heap.list[2])
		}
	})

	t.Run("MinHeap: Testing Pop()", func(t *testing.T) {
		isSmaller := func(a, b int) bool {
			if Compare(a, b) == CR2 {
				return true
			} else {
				return false
			}
		}
		heap := &Heap[int]{[]int{}, isSmaller}
		heap.Push(8)
		heap.Push(6)
		heap.Push(4)

		heap.Pop()
		if heap.Top() == 4 && heap.Top() != 6 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 4)
		}

		heap.Pop()
		if heap.Top() == 6 && heap.Top() != 8 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 6)
		}
		heap.Push(3)
		heap.Push(0)
		heap.Push(-2)
		heap.Push(-5)

		heap.Pop()
		if heap.Top() == -5 && heap.Top() != -2 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), -5)
		}

		heap.Pop()
		if heap.Top() == -2 && heap.Top() != 0 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), -2)
		}

		heap.Push(-20)
		heap.Pop()
		if heap.Top() == -20 && heap.Top() != 3 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), -20)
		}

		heap.Pop()
		if heap.Top() == 6 && heap.Top() != 8 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 6)
		}
	})

	t.Run("MaxHeap: Testing Pop()", func(t *testing.T) {
		isBigger := func(a, b int) bool {
			if Compare(a, b) == CR1 {
				return true
			} else {
				return false
			}
		}
		heap := &Heap[int]{[]int{}, isBigger}
		heap.Push(0)
		heap.Push(1)
		heap.Push(2)
		if heap.Top() != 2 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 2)
		}
		heap.Push(3)
		heap.Push(4)
		heap.Push(5)
		heap.Push(6)
		if heap.Top() != 6 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 6)
		}
		if !(heap.list[1] > heap.list[(1*2)+1] && heap.list[1] > heap.list[(1*2)+2]) {
			t.Error("Index: 1 is not bigger than it's childs")
		}
		if !(heap.list[2] > heap.list[(2*2)+1] && heap.list[2] > heap.list[(2*2)+2]) {
			t.Error("Index 2 is not bigger than it's childs")
		}
	})

	t.Run("MaxHeap: Testing Pop()", func(t *testing.T) {
		isBigger := func(a, b int) bool {
			if Compare(a, b) == 1 {
				return true
			} else {
				return false
			}
		}
		heap := &Heap[int]{[]int{}, isBigger}
		heap.Push(8)
		heap.Push(6)
		heap.Push(4)

		heap.Pop()
		if heap.Top() == 8 && heap.Top() != 6 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 6)
		}

		heap.Pop()
		if heap.Top() == 6 && heap.Top() != 4 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 6)
		}
		heap.Push(3)
		heap.Push(0)
		heap.Push(-2)
		heap.Push(-5)

		heap.Pop()
		if heap.Top() == 4 && heap.Top() != 3 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 4)
		}

		heap.Pop()
		if heap.Top() == 3 && heap.Top() != 0 {
			t.Errorf("Got: %v, Want: %v", heap.Top(), 3)
		}
	})

}
