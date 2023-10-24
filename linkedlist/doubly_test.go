package linkedlist

import (
	"reflect"
	"testing"
)

func TestDoubly(t *testing.T) {
	t.Run("Test AddFirst()", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)

		want := int(2)
		got := list.GetNode(1).(*DoublyNode[int])
		if got.val != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		n1 := list.GetHead()
		val1 := n1.val
		want1 := int(3)
		if val1 != want1 {
			t.Errorf("Got: %v, Want: %v", val1, want1)
		}
		n2 := n1.next
		val2 := n2.val
		want2 := int(2)
		if val2 != want2 {
			t.Errorf("Got: %v, Want: %v", val2, want2)
		}
		want3 := list.GetHead()
		val3 := n2.prev
		if !reflect.DeepEqual(val3, want3) {
			t.Errorf("Got: %v, Want: %v", val3, want3)
		}
		want4 := list.GetTail()
		val4 := n2.next
		if !reflect.DeepEqual(val4, want4) {
			t.Errorf("Got: %v, Want: %v", val4, want4)
		}
	})
	t.Run("Test AddLast()", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)

		want := int(2)
		got := list.GetNode(1).(*DoublyNode[int]).val
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		n1 := list.GetHead()
		val1 := n1.val
		want1 := int(1)
		if val1 != want1 {
			t.Errorf("Got: %v, Want: %v", val1, want1)
		}
		n2 := n1.next
		val2 := n2.val
		want2 := int(2)
		if val2 != want2 {
			t.Errorf("Got: %v, Want: %v", val2, want2)
		}
		want3 := list.GetHead()
		val3 := n2.prev
		if !reflect.DeepEqual(val3, want3) {
			t.Errorf("Got: %v, Want: %v", val3, want3)
		}
		want4 := list.GetTail()
		val4 := n2.next
		if !reflect.DeepEqual(val4, want4) {
			t.Errorf("Got: %v, Want: %v", val4, want4)
		}
	})
	t.Run("Test GetNode()", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)
		list.AddFirst(4)

		want1 := list.GetHead()
		val1 := list.GetNode(0)
		if !reflect.DeepEqual(val1, want1) {
			t.Errorf("Got: %v, Want: %v", val1, want1)
		}

		want2 := want1.next
		val2 := list.GetNode(1)
		if !reflect.DeepEqual(val2, want2) {
			t.Errorf("Got: %v, Want: %v", val2, want2)
		}

		want3 := want2.next
		val3 := list.GetNode(2)
		if !reflect.DeepEqual(val3, want3) {
			t.Errorf("Got: %v, Want: %v", val3, want3)
		}

		want4 := list.GetTail()
		val4 := list.GetNode(3)
		if !reflect.DeepEqual(val4, want4) {
			t.Errorf("Got: %v, Want: %v", val4, want4)
		}
	})

	t.Run("Test DeleteFirst()", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)

		want := int(3)
		got, ok := list.DeleteFirst()
		if !ok {
			t.Error("Not Okay")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want2 := int(2)
		got2, ok2 := list.DeleteFirst()
		if !ok2 {
			t.Error("Not Okay")
		}
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		want3 := int(1)
		got3, ok3 := list.DeleteFirst()
		if !ok3 {
			t.Error("Not Okay")
		}
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}

		got4, ok4 := list.DeleteFirst()
		if ok4 != false {
			t.Error("Not Okay")
		}
		if got4 != 0 {
			t.Errorf("Got: %v, Want: %v", got3, 0)
		}
	})

	t.Run("Test DeleteLast()", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)

		want := int(3)
		got, ok := list.DeleteLast()
		if !ok {
			t.Error("Not Okay")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want2 := int(2)
		got2, ok2 := list.DeleteLast()
		if !ok2 {
			t.Error("Not Okay")
		}
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		want3 := int(1)
		got3, ok3 := list.DeleteLast()
		if !ok3 {
			t.Error("Not Okay")
		}
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}

		got4, ok4 := list.DeleteLast()
		if ok4 != false {
			t.Error("Not Okay")
		}
		if got4 != 0 {
			t.Errorf("Got: %v, Want: %v", got3, 0)
		}
	})
	t.Run("Test DeleteNode()", func(t *testing.T) {
		list := NewDoublyLinkedList[int]()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)

		want := int(2)
		got, ok := list.DeleteNode(1)
		if !ok {
			t.Error("unexpected not-ok")
		}
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

		want2 := int(3)
		got2 := list.Head.next.val

		if got2 != want2 {
			t.Errorf("got: %v, want: %v", got2, want2)
		}

		want3 := int(1)
		got3 := list.Tail.prev.val

		if got3 != want3 {
			t.Errorf("got: %v, want: %v", got3, want3)
		}
	})
}
