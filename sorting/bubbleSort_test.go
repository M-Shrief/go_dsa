package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Testing BubbleSort()", func(t *testing.T) {
		arr := []int{4, 2, 3, 8, 7, 1, 5, 6}

		BubbleSort(arr)
		want := []int{1, 2, 3, 4, 5, 6, 7, 8}
		if !reflect.DeepEqual(arr, want) {
			t.Errorf("Got: %v, Expected: %v", arr, want)
		}

		arr2 := []int{-4, 0, -2, -3, -8, -7, -1, -5, -6}

		BubbleSort(arr2)
		want2 := []int{-8, -7, -6, -5, -4, -3, -2, -1, 0}
		if !reflect.DeepEqual(arr2, want2) {
			t.Errorf("Got: %v, Expected: %v", arr2, want2)
		}
	})
}
