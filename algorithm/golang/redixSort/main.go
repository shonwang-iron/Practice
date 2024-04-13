package main

import (
	"fmt"
)

func main() {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println("Unsorted array:", arr)
	radixSort(arr)
	fmt.Println("Sorted array:", arr)
}

func countingSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}
	fmt.Println("count=>", count)

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	fmt.Println("count<=>", count)

	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

func radixSort(arr []int) {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	for exp := 1; max/exp > 0; exp *= 10 {
		countingSort(arr, exp)
		fmt.Println("arr", arr, "exp", exp)
	}
}
