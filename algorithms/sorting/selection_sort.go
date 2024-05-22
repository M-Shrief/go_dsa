package sorting

import "golang.org/x/exp/constraints"

func SelectionSort[T constraints.Ordered](arr []T, dir Dirction) []T {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {

			if dir == ASC {
				if arr[i] > arr[j] {
					swap(arr, i, j)
				}
			} else if dir == DESC {
				if arr[i] < arr[j] {
					swap(arr, i, j)
				}
			}

		}
	}
	return arr
}
