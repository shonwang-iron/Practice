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

	arr = []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Before merge sorting:", arr)
	arr = mergeSort(arr)
	fmt.Println("After merge sorting:", arr)

	arr = []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Before quick sorting:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("After sorting:", arr)
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

func mergeSort(arr []int) []int {
	n := len(arr)

	// 如果陣列只有一個元素，則返回該元素
	if n <= 1 {
		return arr
	}

	// 將陣列拆分為兩個子陣列，然後分別對它們進行排序
	mid := n / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// 合併排序後的兩個子陣列
	return merge(left, right)
}

func merge(left, right []int) []int {
	nL := len(left)
	nR := len(right)
	result := make([]int, nL+nR)

	// 將兩個已排序的子陣列合併成一個已排序的陣列
	i, j, k := 0, 0, 0
	for i < nL && j < nR {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// 將剩餘的元素添加到已排序的陣列中
	for i < nL {
		result[k] = left[i]
		i++
		k++
	}
	for j < nR {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// 將陣列分區並取得分割點
		p := partition(arr, low, high)

		// 遞歸排序左子陣列
		quickSort(arr, low, p-1)

		// 遞歸排序右子陣列
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// 選擇最後一個元素作為基準值
	pivot := arr[high]

	// i 為小於基準值的元素的最右邊索引
	i := low - 1

	// 將小於基準值的元素移到陣列左邊
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 將基準值移到中間
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// 返回基準值的索引
	return i + 1
}
