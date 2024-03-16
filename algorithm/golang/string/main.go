package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	testCases := []string{"abcde", "hello", "world", "golang"}
	for _, str := range testCases {
		fmt.Printf("String: %s, Non-repeating: %t\n", str, isNonRepeating(str))
	}

	str1 := "listen"
	str2 := "silent"
	fmt.Printf("%s and %s is anagrams: %t\n", str1, str2, isAnagram(str1, str2))

	str3 := "hello"
	str4 := "world"
	fmt.Printf("%s and %s is anagrams: %t\n", str3, str4, isAnagram(str3, str4))

	str := []byte("Mr John Smith    ")
	trueLength := 13 // The actual length is 13, not including the extra space at the end
	fmt.Printf("Original string: %s\n", string(str))
	URLify(str, trueLength)
	fmt.Printf("Replaced string: %s\n", string(str))
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

// Checking anagrams: Given two strings,
// write a method to determine whether one is an anagram of the other.
func isAnagram(str1, str2 string) bool {
	// If the two strings are not equal in length, they are definitely not anagrams.
	if len(str1) != len(str2) {
		return false
	}

	// Convert two strings to character array
	str1Chars := strings.Split(str1, "")
	str2Chars := strings.Split(str2, "")

	// Sort character array
	sort.Strings(str1Chars)
	sort.Strings(str2Chars)

	// Concatenate sorted character arrays into strings
	sortedStr1 := strings.Join(str1Chars, "")
	sortedStr2 := strings.Join(str2Chars, "")

	// Compare sorted strings for equality
	return sortedStr1 == sortedStr2
}

// URLify: Write a method that replaces all spaces in a string with "%20'.
// You can assume that there is enough space at the end of the string to accommodate the extra characters,
// and that you have information about the actual length of the string. Please Use character array
func URLify(str []byte, trueLength int) {
	spaceCount := 0
	// Count the number of spaces
	for i := 0; i < trueLength; i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}

	// Calculate the total length after replacement
	index := trueLength + spaceCount*2

	// Traverse the string from back to front and replace
	for i := trueLength - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str[index-1] = '0'
			str[index-2] = '2'
			str[index-3] = '%'
			index -= 3
		} else {
			str[index-1] = str[i]
			index--
		}
	}
}
