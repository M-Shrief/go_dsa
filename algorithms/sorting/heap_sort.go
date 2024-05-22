package sorting

import "golang.org/x/exp/constraints"

func HeapSort[T constraints.Ordered](arr []T, dir Direction) []T {
	arrSize := len(arr)
	for i := arrSize / 2; i >= 0; i-- {
		heapify(arr, arrSize, i, dir)
	}

	j := arrSize - 1
	for j >= 1 {
		swap(arr, 0, j)
		heapify(arr, j, 0, dir)
		j--
	}

	return arr
}

func heapify[T constraints.Ordered](arr []T, arrSize, i int, dir Direction) {
	largest := i
	leftLeaf := 2*i + 1
	rightLeaf := 2*i + 2

	if dir == ASC {
		if leftLeaf < arrSize && arr[leftLeaf] > arr[largest] {
			largest = leftLeaf
		}
		if rightLeaf < arrSize && arr[rightLeaf] > arr[largest] {
			largest = rightLeaf
		}
	} else if dir == DESC {
		if leftLeaf < arrSize && arr[leftLeaf] < arr[largest] {
			largest = leftLeaf
		}
		if rightLeaf < arrSize && arr[rightLeaf] < arr[largest] {
			largest = rightLeaf
		}
	}

	if largest != i {
		swap(arr, i, largest)
		heapify(arr, arrSize, largest, dir)
	}
}
