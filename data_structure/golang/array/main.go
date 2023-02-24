package main

import "fmt"

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	revArr := reverseArray(arr)
	fmt.Println("reverseArray", revArr)

	queries := [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}
	n := int32(2)

	result := dynamicArray(n, queries)
	fmt.Println("dynamicArray", result) // [7 3]
}

func reverseArray(a []int32) []int32 {
	n := len(a)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func dynamicArray(n int32, queries [][]int32) []int32 {
	seqList := make([][]int32, n)
	lastAns := int32(0)
	results := make([]int32, 0)

	for _, q := range queries {
		t := q[0]
		x := q[1]
		y := q[2]
		idx := (x ^ lastAns) % n

		if t == 1 {
			seqList[idx] = append(seqList[idx], y)
		} else if t == 2 {
			size := int32(len(seqList[idx]))
			lastAns = seqList[idx][y%size]
			results = append(results, lastAns)
		}
	}

	return results
}
