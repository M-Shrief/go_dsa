package array_and_strings

/**
 * Is Unique: Implement an algorithm to determine if a string has all unique characters.
 * What if you cannot use additional data structures?
 *
 */

// version 1 with Data structures.
// Time: O(n), Space: O(n)
// func isUnique_v1(str string) bool {
// 	var mp map[string]bool
// 	mp[string(str[0])] = true
// 	for i := range str {
// 		char := string(str[i])
// 		_, ok := mp[char]
// 		if ok {
// 			return false
// 		} else {
// 			mp[char] = true
// 		}
// 	}
// 	return true
// }

// Version 2, without data structures.
// Time: O(n), Space: O(n)
func isUnique_v2(str string) bool {
	for i := range str {
		if i == len(str) {
			continue
		}
		if str[i] == str[i+1] {
			return false
		}
	}
	return true
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
