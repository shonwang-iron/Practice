package main

import "fmt"

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	revArr := reverseArray(arr)
	fmt.Println(revArr)
}

func reverseArray(a []int32) []int32 {
	n := len(a)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		a[i], a[j] = a[j], a[i]
	}
	return a
}
