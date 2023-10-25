package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Testing Enqueue()", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)

		size := queue.GetSize()
		if size != 3 {
			t.Errorf("Size: %v, Want: %v", size, 3)
		}

		want := int(1)
		n := queue.GetFirst()
		got := n.GetVal()
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want2 := int(2)
		n2 := n.GetNext()
		got2 := n2.GetVal()
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		want3 := int(3)
		n3 := n2.GetNext()
		got3 := n3.GetVal()
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}
	})

	t.Run("Testing Dequeue()", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)

		want := int(1)
		got, ok := queue.Dequeue()
		size := queue.GetSize()
		if size != 2 {
			t.Errorf("Size: %v, Want: %v", size, 2)
		}

		if !ok {
			t.Error("Not Okay")
		}

		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want2 := int(2)
		got2, ok2 := queue.Dequeue()
		size2 := queue.GetSize()
		if size2 != 1 {
			t.Errorf("Size: %v, Want: %v", size2, 1)
		}

		if !ok2 {
			t.Error("Not Okay")
		}

		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		want3 := int(3)
		got3, ok3 := queue.Dequeue()
		size3 := queue.GetSize()
		if size3 != 0 {
			t.Errorf("Size: %v, Want: %v", size3, 0)
		}

		if !ok3 {
			t.Error("Not Okay")
		}

		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}
	})

	t.Run("Testing GetHead()", func(t *testing.T) {
		queue := NewQueue[int]()
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)

		want := int(1)
		n := queue.GetFirst()
		got := n.GetVal()

		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		queue.Dequeue()
		want2 := int(2)
		n2 := queue.GetFirst()
		got2 := n2.GetVal()

		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		queue.Dequeue()
		want3 := int(3)
		n3 := queue.GetFirst()
		got3 := n3.GetVal()

		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}

	})
}
