package main

import "fmt"

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Before insertion sorting:", arr)
	insertionSort(arr)
	fmt.Println("After insertion sorting:", arr)

	arr = []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Before heap sorting:", arr)
	heapSort(arr)
	fmt.Println("After heap sorting:", arr)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func heapify(arr []int, n, i int) {
	largest := i     // 初始化最大值的索引為根節點
	left := 2*i + 1  // 左節點索引
	right := 2*i + 2 // 右節點索引

	// 如果左節點比最大值還大，則將最大值的索引設為左節點
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右節點比最大值還大，則將最大值的索引設為右節點
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值的索引不是根節點，則交換根節點和最大值節點的值，並繼續調整最大堆
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int) {
	n := len(arr)

	// 建立最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 逐一取出堆頂元素，並重新建立最大堆
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}
