package main

import "fmt"

func Add(A, B [][]float64) [][]float64 {
	// run check
	if !CanAddOrSubtract(A, B) {
		fmt.Println("Cannot add the given matrices.")
		return [][]float64{} // return empty matrix
	}

	rows := len(A)
	cols := len(A[0])
	result := make([][]float64, rows)
	for i := range result {
		result[i] = make([]float64, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = A[i][j] + B[i][j]
		}
	}
	return result
}

func Subtract(A, B [][]float64) [][]float64 {
	// run check
	if !CanAddOrSubtract(A, B) {
		fmt.Println("Cannot subtract the given matrices.")
		return [][]float64{} // return empty matrix
	}

	rows := len(A)
	cols := len(A[0])
	result := make([][]float64, rows)
	for i := range result {
		result[i] = make([]float64, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = A[i][j] - B[i][j]
		}
	}
	return result
}

func Multiply(A, B [][]float64) [][]float64 {
	// run check
	if !CanMultiply(A, B) {
		fmt.Println("Cannot multiply the given matrices.")
		return [][]float64{} // return empty matrix
	}

	// multiplication using dynamic dimensions
	rowsA := len(A)
	colsA := len(A[0])
	colsB := len(B[0])

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			sum := 0.0
			for k := 0; k < colsA; k++ {
				sum += A[i][k] * B[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

func Transpose(A [][]float64) [][]float64 {
	// run valid matrix check
	if !IsValidMatrix(A) {
		fmt.Println("Cannot transpose the given matrices.")
		return [][]float64{} // return empty matrix
	}

	rows := len(A)
	cols := len(A[0])

	result := make([][]float64, cols)
	for i := range result {
		result[i] = make([]float64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = A[i][j]
		}
	}
	return result
}
