package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 8, 4, 7}
	target := 8
	index := linearSearch(arr, target)
	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found\n", target)
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
