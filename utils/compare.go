package utils

import "golang.org/x/exp/constraints"

/*
Compare any 2 paramters type that supports the operators < <= >= >.

- Returns 1, if a > b.

- Returns -1, if a < b.

- Returns 0, if a == b.

Note: Not fit to compare pointers.
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
