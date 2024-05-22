package sorting

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	t.Run("Testing ASC sorting", func(t *testing.T) {
		arr := []int{1, 5, 3, 7, 4, 47, 77}

		got := HeapSort(arr, ASC)
		want := []int{1, 3, 4, 5, 7, 47, 77}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Testing DESC sorting", func(t *testing.T) {
		arr := []int{1, 5, 3, 7, 4, 47, 77}

		got := HeapSort(arr, DESC)
		want := []int{77, 47, 7, 5, 4, 3, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})
}
