package tree

import (
	"reflect"
	"testing"
)

func TestBST(t *testing.T) {
	t.Run("Test Insert()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)

		want := int(5)
		n := tree.GetRoot()
		got := n.GetVal()
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want1 := int(2)
		n1 := n.left
		got1 := n1.GetVal()
		if got1 != want1 {
			t.Errorf("Got: %v, Want: %v", got1, want1)
		}

		want2 := int(1)
		n2 := n1.left
		got2 := n2.GetVal()
		if got1 != want1 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		want3 := int(3)
		n3 := n1.right
		got3 := n3.GetVal()
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		want4 := int(7)
		n4 := n.right
		got4 := n4.GetVal()
		if got1 != want1 {
			t.Errorf("Got: %v, Want: %v", got4, want4)
		}

		want5 := int(6)
		n5 := n4.left
		got5 := n5.GetVal()
		if got5 != want5 {
			t.Errorf("Got: %v, Want: %v", got5, want5)
		}

		want6 := int(8)
		n6 := n4.right
		got6 := n6.GetVal()
		if got6 != want6 {
			t.Errorf("Got: %v, Want: %v", got6, want6)
		}
	})

	t.Run("Test Search()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(6)

		want := int(5)
		n, ok := tree.Search(5)
		got := n.GetVal()

		if !ok {
			t.Error("Not Okay")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want1 := int(2)
		n1, ok1 := tree.Search(2)
		got1 := n1.GetVal()

		if !ok1 {
			t.Error("Not Okay")
		}
		if got1 != want1 {
			t.Errorf("Got: %v, Want: %v", got1, want1)
		}

		want2 := int(1)
		n2, ok2 := tree.Search(1)
		got2 := n2.GetVal()
		if !ok2 {
			t.Error("Not Okay")
		}
		if got2 != want2 {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		want3 := int(3)
		n3, ok3 := tree.Search(3)
		got3 := n3.GetVal()
		if !ok3 {
			t.Error("Not Okay")
		}
		if got3 != want3 {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}

		want4 := int(6)
		n4, ok4 := tree.Search(6)
		got4 := n4.GetVal()
		if !ok4 {
			t.Error("Not Okay")
		}
		if got4 != want4 {
			t.Errorf("Got: %v, Want: %v", got4, want4)
		}
	})

	t.Run("Test minimunVal()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		want := int(1)
		got, ok := tree.Minimum()

		if !ok {
			t.Error("Not Okay")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Test minimunVal()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		want := int(8)
		got, ok := tree.Maximum()
		if !ok {
			t.Error("Not Okay")
		}
		if got != want {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})
	t.Run("Test GetParent()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		want := tree.root
		got := tree.GetParent(2)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}

		want1 := tree.root.left
		got1 := tree.GetParent(1)
		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("Got: %v, Want: %v", got1, want1)
		}
		got2 := tree.GetParent(3)
		if !reflect.DeepEqual(got2, want1) {
			t.Errorf("Got: %v, Want: %v", got2, want1)
		}

		want2 := tree.root.right
		got3 := tree.GetParent(6)
		if !reflect.DeepEqual(got3, want2) {
			t.Errorf("Got: %v, Want: %v", got3, want2)
		}
		got4 := tree.GetParent(8)
		if !reflect.DeepEqual(got4, want2) {
			t.Errorf("Got: %v, Want: %v", got4, want2)
		}
	})

	t.Run("Test Remove()", func(t *testing.T) {
		tree := NewBST[int]()
		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		tree.Remove(2)
		_, ok := tree.Search(2)
		if ok != false {
			t.Error("Not Okay")
		}
		replacement := tree.GetRoot().left
		if replacement.val != 1 {
			t.Errorf("Got: %v, Want: %v", replacement, 1)
		}
		if replacement.right.val != 3 {
			t.Errorf("Got: %v, Want: %v", replacement, 3)
		}
		if tree.size != 6 {
			t.Errorf("Size: %v, want: %v", tree.size, 6)
		}

		tree.Remove(1)
		_, ok3 := tree.Search(1)
		if ok3 != false {
			t.Error("Not Okay")
		}
		replacement3 := tree.GetRoot().left
		if replacement3.val != 3 {
			t.Errorf("Got: %v, Want: %v", replacement3.val, 3)
		}
		if replacement3.left != nil {
			t.Error("Replacement left is not nil")
		}
		if replacement3.right != nil {
			t.Error("Replacement right is not nil")
		}
		if tree.size != 5 {
			t.Errorf("Size: %v, want: %v", tree.size, 5)
		}

		tree.Remove(100)

		tree.Remove(3)
		_, ok4 := tree.Search(3)
		if ok4 != false {
			t.Error("Not Okay")
		}
		replacement4 := tree.GetRoot().left
		if replacement4 != nil {
			t.Error("Replacement is not nil")
		}
		if tree.size != 4 {
			t.Errorf("Size: %v, want: %v", tree.size, 4)
		}

		tree.Remove(7)
		_, ok2 := tree.Search(7)
		if ok2 != false {
			t.Error("Not Okay")
		}

		replacement2 := tree.GetRoot().right
		if replacement2.val != 6 {
			t.Errorf("Got: %v, Want: %v", replacement2.val, 6)
		}
		if replacement2.left != nil {
			t.Errorf("Got: %v, Want: %v", replacement2, 8)
		}

		tree.Remove(8)
		_, ok5 := tree.Search(8)
		if ok5 != false {
			t.Error("Not Okay")
		}
		replacement6 := tree.GetRoot().right.right
		if replacement6 != nil {
			t.Error("Replacement is not nil")
		}

		tree.Remove(6)
		_, ok6 := tree.Search(6)
		if ok6 != false {
			t.Error("Not Okay")
		}

		replacement7 := tree.GetRoot().right
		if replacement7 != nil {
			t.Error("Replacement is not nil")
		}

		tree.Remove(5)
		_, ok7 := tree.Search(5)
		if ok7 != false {
			t.Error("Not Okay")
		}
		if tree.root != nil {
			t.Error("Root is not Nil")
		}

	})

	t.Run("Test Breadth First Traversal, BFT()", func(t *testing.T) {
		tree := NewBST[int]()

		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		// Tree looks like [5,2,7,1,3,6,8]
		/*
					5
				2		7
			  1	  3	  6    8
		*/
		got := tree.BFT()
		want := make([]int, tree.GetSize())
		want = append(want, 5, 2, 7, 1, 3, 6, 8)

		if len(got) != len(want) {
			t.Errorf("Got: %v. Want: %v", len(got), len(want))
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Test Breadth First Search, BFS()", func(t *testing.T) {
		tree := NewBST[int]()

		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		// Tree looks like [5,2,7,1,3,6,8]
		/*
					5
				2		7
			  1	  3	  6    8
		*/
		got, err := tree.BFS(2)
		two := tree.GetRoot().left
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if got.val != two.val {
			t.Errorf("Got: %v, Want: %v", got.val, two.val)
		}
		if got.left != two.left {
			t.Errorf("Got: %v, Want: %v", got.left, two.left)
		}
		if got.right != two.right {
			t.Errorf("Got: %v, Want: %v", got.right, two.right)
		}

		got2, err := tree.BFS(8)
		eight := tree.GetRoot().right.right
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if got2.val != eight.val {
			t.Errorf("Got: %v, Want: %v", got2.val, eight.val)
		}
		if got2.left != eight.left {
			t.Errorf("Got: %v, Want: %v", got2.left, eight.left)
		}
		if got2.right != eight.right {
			t.Errorf("Got: %v, Want: %v", got2.right, eight.right)
		}
	})

	t.Run("Test Depth First Traversal, DFT()", func(t *testing.T) {
		tree := NewBST[int]()

		tree.Insert(5)
		tree.Insert(2)
		tree.Insert(1)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(6)
		tree.Insert(8)

		/*
			Tree Looks like:
						5
					2		7
				  1	  3	  6    8
		*/
		got1 := tree.DFT(DFTPreOrder)
		want1 := make([]int, tree.GetSize())
		want1 = append(want1, 5, 2, 1, 3, 7, 6, 8)

		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("Got: %v, Want: %v", got1, want1)
		}

		got2 := tree.DFT(DFTInOrder)
		want2 := make([]int, tree.GetSize())
		want2 = append(want2, 1, 2, 3, 5, 6, 7, 8)

		if !reflect.DeepEqual(got2, want2) {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		got3 := tree.DFT(DFTPostOrder)
		want3 := make([]int, tree.GetSize())
		want3 = append(want3, 1, 3, 2, 6, 8, 7, 5)

		if !reflect.DeepEqual(got3, want3) {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}
	})
}
