package main

import "fmt"

func main() {
	image := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("The original image：")
	printMatrix(image)
	rotateImage(image)
	fmt.Println("\nImage rotated 90 degrees：")
	printMatrix(image)

	fmt.Println("================================================================")

	matrix := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	fmt.Println("original matrix：")
	printMatrix(matrix)

	setZeroes(matrix)

	fmt.Println("\nThe matrix after setting it to zero：")
	printMatrix(matrix)
}

// print matrix function
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// Rotation matrix:
// Given an N*N matrix representing an image,
// where each pixel in the image is 4 bits,
// please write a method to rotate the image 90 degrees.
// You can do it in-place in the same memory ( in place) completed?
func rotateImage(matrix [][]int) {
	n := len(matrix)

	// Flip diagonally first
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// flip horizontally
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

// Zero matrix: Please write this algorithm.
// If an element in the M*N matrix is ​​0, set the entire column and column where the element is located to 0.
func setZeroes(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])

	// Use two arrays to record which rows and columns need to be set to zero
	rowZero := make([]bool, rows)
	colZero := make([]bool, cols)

	// Traverse the matrix and record the rows and columns with zero
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				rowZero[i] = true
				colZero[j] = true
			}
		}
	}

	// Set the rows and columns that need to be zeroed to zero
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rowZero[i] || colZero[j] {
				matrix[i][j] = 0
			}
		}
	}
}
