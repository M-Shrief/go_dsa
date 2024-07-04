package utils

import "golang.org/x/exp/constraints"

/*
Compare 2 paramters that can be ordered.

Returns 1, if a > b.

Returns -1, if a < b.

Returns 0, if a == b.
*/
func Compare[T constraints.Ordered](a, b T) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
