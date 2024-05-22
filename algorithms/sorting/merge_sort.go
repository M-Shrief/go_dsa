package sorting

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](arr []T, dir Direction) []T {
	if len(arr) < 2 {
		return arr
	}
	first := MergeSort(arr[:len(arr)/2], dir)
	second := MergeSort(arr[len(arr)/2:], dir)
	return merge(first, second, dir)
}

func merge[T constraints.Ordered](leftArr, rightArr []T, dir Direction) []T {
	final := []T{}
	i := 0
	j := 0
	for i < len(leftArr) && j < len(rightArr) {
		if (dir == ASC && leftArr[i] < rightArr[j]) || (dir == DESC && leftArr[i] > rightArr[j]) {
			final = append(final, leftArr[i])
			i++
		} else {
			final = append(final, rightArr[j])
			j++
		}
	}
	for ; i < len(leftArr); i++ {
		final = append(final, leftArr[i])
	}
	for ; j < len(rightArr); j++ {
		final = append(final, rightArr[j])
	}
	return final
}
