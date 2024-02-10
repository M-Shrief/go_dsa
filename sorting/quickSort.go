package sorting

import "golang.org/x/exp/constraints"

func QuickSort[T constraints.Ordered](arr []T) []T {
	quickSortRange[T](arr, 0, len(arr)-1)
	return arr
}

func quickSortRange[T constraints.Ordered](arr []T, low, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		// pIndex is the partitioning index, arr[pIndex] is now at the right place
		pIndex := partition[T](arr, low, high)
		quickSortRange[T](arr, low, pIndex-1)
		quickSortRange[T](arr, pIndex+1, high)
	}
}

func partition[T constraints.Ordered](arr []T, low, high int) int {
	pivot := arr[high]
	index := low - 1

	for i := low; i < high; i++ {
		if arr[i] < pivot {
			index++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}
