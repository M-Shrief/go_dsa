package graph

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	t.Run("Testing BFT()", func(t *testing.T) {
		g := NewGraph[int]()

		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 2)
		g.AddEdge(2, 0)
		g.AddEdge(2, 3)
		g.AddEdge(3, 3)
		g.AddEdge(3, 4)
		g.AddEdge(5, 4)

		got1, ok := g.BFT(0)
		want1 := []int{0, 1, 2, 3, 4}
		if ok != nil {
			t.Error("Not Okay")
		}
		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("Got: %v, Want: %v", got1, want1)
		}

		got2, ok := g.BFT(1)
		want2 := []int{1, 2, 0, 3, 4}
		if ok != nil {
			t.Error("Not Okay")
		}
		if !reflect.DeepEqual(got2, want2) {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		got3, ok := g.BFT(1)
		want3 := []int{1, 2, 0, 3, 4}
		if ok != nil {
			t.Error("Not Okay")
		}
		if !reflect.DeepEqual(got3, want3) {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}

		got4, ok := g.BFT(4)
		want4 := []int{4}
		if ok != nil {
			t.Error("Not Okay")
		}
		if !reflect.DeepEqual(got4, want4) {
			t.Errorf("Got: %v, Want: %v", got4, want4)
		}
	})

	t.Run("Testing DFT", func(t *testing.T) {
		g := NewGraph[int]()

		g.AddEdge(0, 1)
		g.AddEdge(0, 2)
		g.AddEdge(1, 3)
		g.AddEdge(2, 0)
		g.AddEdge(2, 3)
		g.AddEdge(3, 3)
		g.AddEdge(3, 4)
		g.AddEdge(4, 5)

		got1, ok := g.DFT(0)
		want1 := []int{0, 1, 3, 4, 5, 2}
		if ok != nil {
			t.Error("Not Okay")
		}
		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("Got: %v, Want: %v", got1, want1)
		}

		got2, ok := g.DFT(1)
		want2 := []int{1, 3, 4, 5}
		if ok != nil {
			t.Error("Not Okay")
		}
		if !reflect.DeepEqual(got2, want2) {
			t.Errorf("Got: %v, Want: %v", got2, want2)
		}

		got3, ok := g.DFT(2)
		want3 := []int{2, 0, 1, 3, 4, 5}
		if ok != nil {
			t.Error("Not Okay")
		}
		if !reflect.DeepEqual(got3, want3) {
			t.Errorf("Got: %v, Want: %v", got3, want3)
		}

		got4, ok := g.DFT(3)
		want4 := []int{3, 4, 5}
		if ok != nil {
			t.Error("Not Okay")
		}
		if !reflect.DeepEqual(got4, want4) {
			t.Errorf("Got: %v, Want: %v", got4, want4)
		}
	})
}
