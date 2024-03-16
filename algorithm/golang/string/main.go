package main

import "fmt"

func main() {
	testCases := []string{"abcde", "hello", "world", "golang"}

	for _, str := range testCases {
		fmt.Printf("String: %s, Non-repeating: %t\n", str, isNonRepeating(str))
	}
}

// Non-repeating: Implement an algorithm to
// whether the word sources in a string are non-repeating.
// What if other data structures cannot be used?
func isNonRepeating(str string) bool {
	// Assume the string contains only ASCII characters (0-127)
	if len(str) > 128 {
		// If the string length is greater than the number of ASCII characters,
		// there must be duplicate characters
		return false
	}

	var charSet [128]bool
	for _, char := range str {
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}
	return true
}
