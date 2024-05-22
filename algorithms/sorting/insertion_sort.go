package sorting

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](arr []T, dir Dirction) []T {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		if dir == ASC {
			for j >= 0 && arr[j] > key {
				arr[j+1] = arr[j]
				j--
			}
		} else if dir == DESC {
			for j >= 0 && arr[j] < key {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
	}
	return arr
}
