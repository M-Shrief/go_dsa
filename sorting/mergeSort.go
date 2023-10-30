package sorting

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](arr []T) []T {
	arrLength := len(arr)
	if arrLength <= 1 {
		return arr
	}

	halfIndex := arrLength / 2
	mergedHavles := merge(
		MergeSort[T](arr[0:halfIndex]),
		MergeSort[T](arr[halfIndex:]),
	)
	return mergedHavles
}

func merge[T constraints.Ordered](leftArr []T, rightArr []T) []T {
	var mergeArr []T
	i, j := 0, 0
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] < rightArr[j] {
			mergeArr = append(mergeArr, leftArr[i])
			i++
		} else {
			mergeArr = append(mergeArr, rightArr[j])
			j++
		}
	}
	mergedArr := append(mergeArr, leftArr[i:]...)
	mergedArr = append(mergedArr, rightArr[j:]...)
	return mergedArr

}
