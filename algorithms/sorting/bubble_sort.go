package sorting

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](arr []T, dir Dirction) []T {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if dir == ASC {
				if arr[i] < arr[j] {
					swap(arr, i, j)
				}
			} else if dir == DESC {
				if arr[i] > arr[j] {
					swap(arr, i, j)
				}
			}
		}
	}

	return arr
}
