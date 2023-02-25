package main

import (
	"fmt"
	"math"
)

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
	fmt.Println("dynamicArray", result)

	maxSum := hourglassSum([][]int32{
		{0, -4, -6, 0, -7, -6},
		{-1, -2, -6, -8, -3, -1},
		{-8, -4, -2, -8, -8, -6},
		{-3, -1, -2, -5, -7, -4},
		{-3, -5, -3, -6, -6, -6},
		{-3, -6, 0, -8, -6, -7},
	})
	fmt.Println("hourglassSum", maxSum)

	result64 := arrayManipulation(int32(5),
		[][]int32{
			{1, 2, 100},
			{2, 5, 100},
			{3, 4, 100},
		})
	fmt.Println("arrayManipulation", result64)

	result32 := rotateLeft(4, []int32{1, 2, 3, 4, 5})
	fmt.Println("rotateLeft", result32)

	result32 = matchingStrings(
		[]string{
			"aba",
			"baba",
			"aba",
			"xzxb",
		},
		[]string{
			"aba",
			"xzxb",
			"ab",
		},
	)
	fmt.Println("matchingStrings", result32)
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

func hourglassSum(arr [][]int32) int32 {
	maxSum := int32(math.MinInt32)

	for i := 0; i <= len(arr)-3; i++ {
		for j := 0; j <= len(arr[0])-3; j++ {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n+1)
	for _, q := range queries {
		a := q[0]
		b := q[1]
		k := q[2]
		arr[a-1] += int64(k)
		arr[b] -= int64(k)
	}

	var max int64 = 0
	var sum int64 = 0
	for _, val := range arr {
		sum += val
		if sum > max {
			max = sum
		}
	}
	return max
}

func rotateLeft(d int32, arr []int32) []int32 {
	return append(arr[d:], arr[:d]...)
}

func matchingStrings(stringList []string, queries []string) []int32 {
	freqMap := make(map[string]int32)
	for _, s := range stringList {
		freqMap[s]++
	}

	res := make([]int32, len(queries))
	for i, q := range queries {
		res[i] = freqMap[q]
	}

	return res
}
