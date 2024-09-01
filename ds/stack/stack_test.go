package stack

import "testing"

func TestStack(t *testing.T) {
	t.Run("Testing Push()", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		want := int(3)
		got := stack.GetTop()
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want2 := int(2)
		n := stack.list.GetHead().GetNext()
		got2 := n.GetVal()
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		want3 := int(1)
		n2 := n.GetNext()
		got3 := n2.GetVal()
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}

		size := stack.GetSize()
		if size != 3 {
			t.Errorf("Size: %v, Want: %v", size, 3)
		}
	})

	t.Run("Testing Pop()", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		want := int(3)
		got, ok := stack.Pop()
		if !ok {
			t.Error("Not Okay")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
		size := stack.GetSize()
		if size != 2 {
			t.Errorf("Size: %v, Want: %v", size, 2)
		}

		want2 := int(2)
		got2, ok2 := stack.Pop()
		if !ok2 {
			t.Error("Not Okay")
		}
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}
		size2 := stack.GetSize()
		if size2 != 1 {
			t.Errorf("Size: %v, Want: %v", size, 1)
		}

		want3 := int(1)
		got3, ok3 := stack.Pop()
		if !ok3 {
			t.Error("Not Okay")
		}
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}
		size3 := stack.GetSize()
		if size3 != 0 {
			t.Errorf("Size: %v, Want: %v", size, 0)
		}
	})

}
