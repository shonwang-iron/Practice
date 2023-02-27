package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 8, 4, 7}
	target := 8
	index := linearSearch(arr, target)
	if index != -1 {
		fmt.Printf("Linear search: Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Linear search: Target %d not found\n", target)
	}

	arr = []int{1, 2, 4, 5, 7, 8}
	target = 7
	index = binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Binary search: Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Binary search: Target %d not found\n", target)
	}
}

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
