package linkedlist

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {

	t.Run("Test AddFirst()", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)

		want := int(2)
		got := list.GetNode(1)

		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Test AddFirst()", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)

		want := []int{3, 2, 1}
		got := []int{}

		current := list.Head
		got = append(got, current.Val)

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})

	t.Run("Test AddLast()", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)
		list.AddLast(4)

		want := []int{3, 2, 1, 4}
		got := []int{}

		current := list.Head
		got = append(got, current.Val)

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Test DeleteFirst()", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)
		list.AddLast(4)

		want := int(3)
		got, ok := list.DeleteFirst()

		if !ok {
			t.Error("Not Okay")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Test DeleteLast()", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)
		list.AddLast(4)

		want := int(4)
		got, ok := list.DeleteLast()

		if !ok {
			t.Error("Not Okay!")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Test DeleteNode(pos)", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)

		want := int(2)
		got, ok := list.DeleteNode(1)

		if !ok {
			t.Error("Not Okay!")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want2 := 1
		got2 := list.Head.Next.Val
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}
	})

	t.Run("Test Reverse()", func(t *testing.T) {
		list := NewSinglyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)

		want := []int{1, 2, 3}
		got := []int{}

		list.Reverse()
		current := list.Head
		got = append(got, current.Val)

		for current.Next != nil {
			current = current.Next
			got = append(got, current.Val)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})
}
