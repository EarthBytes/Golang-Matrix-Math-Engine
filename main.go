package main

import "fmt"

func main() {
	A := [][]float64{
		{1, 2},
		{3, 4},
	}

	B := [][]float64{
		{5, 6},
		{7, 8},
	}

	result := Multiply(A, B)

	fmt.Println(result)
}

func CanMultiply(A [][]float64, B [][]float64) bool {
	if len(A) == 0 || len(B) == 0 {
		return false
	}

	// ensure matrix A is a proper rectangle
	colsA := len(A[0])
	for _, row := range A {
		if len(row) != colsA {
			return false // found jagged matrix
		}
	}

	// ensure matirx B is a proper rectangle
	colsB := len(B[0])
	for _, row := range B {
		if len(row) != colsB {
			return false // found jagged matrix
		}
	}

	// check standard matrix multiplication
	return colsA == len(B)
}

func Multiply(A [][]float64, B [][]float64) [][]float64 {

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
