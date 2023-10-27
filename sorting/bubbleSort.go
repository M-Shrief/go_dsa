package sorting

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[i] {
				swap[T](arr, j, i)
			}
		}
	}
	return arr
}

func swap[T constraints.Ordered](arr []T, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
