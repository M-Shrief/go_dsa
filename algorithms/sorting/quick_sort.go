package sorting

import "golang.org/x/exp/constraints"

// QuickSort algorithm
func QuickSort[T constraints.Ordered](arr []T, dir Direction, leftmostIndex, rightmostIndex int) []T {
	if leftmostIndex < rightmostIndex {
		partitionIndex := partition(arr, dir, leftmostIndex, rightmostIndex)

		QuickSort(arr, dir, leftmostIndex, partitionIndex-1)
		QuickSort(arr, dir, partitionIndex+1, rightmostIndex)
	}
	return arr
}

func partition[T constraints.Ordered](arr []T, dir Direction, leftmostIndex, rightmostIndex int) int {
	pivot := arr[rightmostIndex]
	storeIndex := leftmostIndex - 1

	if dir == ASC {
		for j := leftmostIndex; j <= rightmostIndex-1; j++ {
			if arr[j] < pivot {
				storeIndex++
				swap(arr, storeIndex, j)
			}
		}
	} else if dir == DESC {
		for j := leftmostIndex; j <= rightmostIndex-1; j++ {
			if arr[j] > pivot {
				storeIndex++
				swap(arr, storeIndex, j)
			}
		}
	}

	swap(arr, storeIndex+1, rightmostIndex)
	return storeIndex + 1
}
