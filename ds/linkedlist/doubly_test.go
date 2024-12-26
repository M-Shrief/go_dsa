package linkedlist

import (
	"reflect"
	"testing"
)

func TestDoubly(t *testing.T) {
	t.Run("Test AddFirst()", func(t *testing.T) {
		list := NewDoubly[int]()
		list.AddFirst(1)
		list.AddFirst(2)
		list.AddFirst(3)

		want := int(2)
		got := list.GetNode(1)
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

		size := list.size
		if size != 3 {
			t.Errorf("Size: %v, Want: %v", size, 3)
		}
	})
	t.Run("Test AddLast()", func(t *testing.T) {
		list := NewDoubly[int]()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)

		want := int(2)
		got := list.GetNode(1)
		if got.val != want {
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
		size := list.size
		if size != 3 {
			t.Errorf("Size: %v, Want: %v", size, 3)
		}
	})
	t.Run("Test GetNode()", func(t *testing.T) {
		list := NewDoubly[int]()
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
		list := NewDoubly[int]()
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
		size := list.size
		if size != 2 {
			t.Errorf("Size: %v, Want: %v", size, 2)
		}

		want2 := int(2)
		got2, ok2 := list.DeleteFirst()
		if !ok2 {
			t.Error("Not Okay")
		}
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}
		size2 := list.size
		if size2 != 1 {
			t.Errorf("Size: %v, Want: %v", size2, 1)
		}

		want3 := int(1)
		got3, ok3 := list.DeleteFirst()
		if !ok3 {
			t.Error("Not Okay")
		}
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}
		size3 := list.size
		if size3 != 0 {
			t.Errorf("Size: %v, Want: %v", size, 0)
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
		list := NewDoubly[int]()
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
		size := list.size
		if size != 2 {
			t.Errorf("Size: %v, Want: %v", size, 2)
		}

		want2 := int(2)
		got2, ok2 := list.DeleteLast()
		if !ok2 {
			t.Error("Not Okay")
		}
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}
		size2 := list.size
		if size2 != 1 {
			t.Errorf("Size: %v, Want: %v", size2, 1)
		}

		want3 := int(1)
		got3, ok3 := list.DeleteLast()
		if !ok3 {
			t.Error("Not Okay")
		}
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}
		size3 := list.size
		if size3 != 0 {
			t.Errorf("Size: %v, Want: %v", size3, 0)
		}

		got4, ok4 := list.DeleteLast()
		if ok4 != false {
			t.Error("Not Okay")
		}
		if got4 != 0 {
			t.Errorf("Got: %v, Want: %v", got3, 0)
		}
	})
	t.Run("Test DeleteByPosition()", func(t *testing.T) {
		list := NewDoubly[int]()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)

		want := int(2)
		got, ok := list.DeleteByPosition(1)
		if !ok {
			t.Error("unexpected not-ok")
		}
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

		want2 := int(3)
		got2 := list.head.next.val

		if got2 != want2 {
			t.Errorf("got: %v, want: %v", got2, want2)
		}

		want3 := int(1)
		got3 := list.tail.prev.val

		if got3 != want3 {
			t.Errorf("got: %v, want: %v", got3, want3)
		}
	})

	t.Run("Test DeleteByNode()", func(t *testing.T) {
		list := NewDoubly[int]()
		list.AddLast(1)
		list.AddLast(2)
		list.AddLast(3)
		list.AddLast(4)

		want := int(1)
		got, _ := list.DeleteByNode(list.head)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
		if list.size != 3 {
			t.Errorf("Size: %v, want: %v", list.size, 3)
		}

		want2 := int(3)
		got2, _ := list.DeleteByNode(list.head.next)
		if got2 != want2 {
			t.Errorf("got: %v, want: %v", got2, want2)
		}

		if list.size != 2 {
			t.Errorf("Size: %v, want: %v", list.size, 1)
		}

		want3 := int(4)
		got3, _ := list.DeleteByNode(list.tail)
		if got3 != want3 {
			t.Errorf("got: %v, want: %v", got3, want3)
		}

		if list.size != 1 {
			t.Errorf("Size: %v, want: %v", list.size, 1)
		}

		want4 := int(2)
		got4, _ := list.DeleteByNode(list.head)
		if got4 != want4 {
			t.Errorf("got: %v, want: %v", got4, want4)
		}

		if list.size != 0 {
			t.Errorf("Size: %v, want: %v", list.size, 0)
		}

		_, isDeleted := list.DeleteByNode(nil)
		if isDeleted {
			t.Errorf("Not Okay!")
		}
	})

}
