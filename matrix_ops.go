package main

import "sync"

func Add(A, B [][]float64) ([][]float64, error) {
	// run check
	if !CanAddOrSubtract(A, B) {
		return nil, ErrCannotAdd // return error for addition if check fails
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
	return result, nil // return the result and nil error if successful
}

func Subtract(A, B [][]float64) ([][]float64, error) { // return error for subtraction if check fails
	// run check
	if !CanAddOrSubtract(A, B) {
		return nil, ErrCannotSubtract
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
	return result, nil
}

func Multiply(A, B [][]float64) ([][]float64, error) {
	// run check
	if !CanMultiply(A, B) {
		return nil, ErrCannotMultiply
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
	return result, nil
}

// MultiplyParallel does the same as Multiply but uses one goroutine per row of the result
func MultiplyParallel(A, B [][]float64) ([][]float64, error) {
	// run check (same as Multiply)
	if !CanMultiply(A, B) {
		return nil, ErrCannotMultiply
	}

	rowsA := len(A)
	colsA := len(A[0])
	colsB := len(B[0])

	result := make([][]float64, rowsA)

	var wg sync.WaitGroup
	wg.Add(rowsA)

	// each goroutine computes one row of the result
	for i := 0; i < rowsA; i++ {
		go func(i int) {
			defer wg.Done()

			row := make([]float64, colsB)
			for j := 0; j < colsB; j++ {
				sum := 0.0
				for k := 0; k < colsA; k++ {
					sum += A[i][k] * B[k][j]
				}
				row[j] = sum
			}
			result[i] = row // each goroutine writes only its own row
		}(i)
	}

	wg.Wait()
	return result, nil
}

func Transpose(A [][]float64) ([][]float64, error) {
	// run valid matrix check
	if !IsValidMatrix(A) {
		return nil, ErrCannotTranspose
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
	return result, nil
}
