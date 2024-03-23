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

	fmt.Println("================================================================")
	str1 := "listen"
	str2 := "silent"
	fmt.Printf("%s and %s is anagrams: %t\n", str1, str2, isAnagram(str1, str2))

	str3 := "hello"
	str4 := "world"
	fmt.Printf("%s and %s is anagrams: %t\n", str3, str4, isAnagram(str3, str4))

	fmt.Println("================================================================")
	str := []byte("Mr John Smith    ")
	trueLength := 13 // The actual length is 13, not including the extra space at the end
	fmt.Printf("Original string: %s\n", string(str))
	URLify(str, trueLength)
	fmt.Printf("Replaced string: %s\n", string(str))

	fmt.Println("================================================================")
	testCases = []string{"abcba", "aabbc", "tactcoapapa", "Tact Coa", "A man a plan a canal Panama", "hello world"}
	for _, str := range testCases {
		fmt.Printf("String: %s, Is Palindrome Anagram: %t\n", str, isPermutationOfPalindrome(str))
	}

	fmt.Println("================================================================")
	testCases2 := [][]string{
		{"pale", "ple"},
		{"pales", "pale"},
		{"pale", "bale"},
		{"pale", "bake"},
	}
	for _, tc := range testCases2 {
		fmt.Printf("String 1: %s, String 2: %s, Is One or Zero Edit Away: %t\n", tc[0], tc[1], isOneOrZeroEditAway(tc[0], tc[1]))
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

// Palindrome anagram: You get a string, please write a function to check whether it is a palindrome anagram.
// A palindrome means a word or phrase that is the same whether it is read forward or backward.
// Anagrams are letters rearranged. Palindromes are not limited to dictionary words.
func isPermutationOfPalindrome(phrase string) bool {
	table := buildCharFrequencyTable(phrase)
	return checkMaxOneOdd(table)
}

func checkMaxOneOdd(table []int) bool {
	foundOdd := false
	for _, count := range table {
		if count%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}

func getCharNumber(c rune) int {
	a := int('a')
	z := int('z')
	A := int('A')
	Z := int('Z')
	value := int(c)
	if a <= value && value <= z {
		return value - a
	} else if A <= value && value <= Z {
		return value - A
	}
	return -1
}

func buildCharFrequencyTable(phrase string) []int {
	table := make([]int, int('z')-int('a')+1)
	for _, char := range phrase {
		x := getCharNumber(char)
		if x != -1 {
			table[x]++
		}
	}
	return table
}

func isOneOrZeroEditAway(str1, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)

	// Calculate the length difference of two strings
	diff := len1 - len2
	if diff < 0 {
		diff = -diff
	}

	// If the length difference is greater than 1, the edit distance must be greater than 1
	if diff > 1 {
		return false
	}

	// For a length difference of 1, after finding the first different character position, compare the subsequent characters.
	i, j := 0, 0
	edits := 0
	for i < len1 && j < len2 {
		if str1[i] != str2[j] {
			edits++
			if edits > 1 {
				return false
			}

			// For the case where the length differs by 1, align different character positions
			if diff == 1 {
				if len1 > len2 {
					i++
				} else {
					j++
				}
				continue
			}
		}

		i++
		j++
	}

	// If the lengths of the two strings are different and one of the strings has been traversed, the remaining part of the other string can only be inserted.
	if i < len1 || j < len2 {
		edits++
	}
	return edits <= 1
}
