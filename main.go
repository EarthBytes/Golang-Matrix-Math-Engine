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

	result := make([][]float64, 2)
	for i := range result {
		result[i] = make([]float64, 2)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {

			sum := 0.0

			for k := 0; k < 2; k++ {
				sum += A[i][k] * B[k][j]
			}

			result[i][j] = sum
		}
	}
	fmt.Println(result)
}
