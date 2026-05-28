package main

import "fmt"

// LinearForward computes y = xW + b
// x is the input, W is the weight matrix, and b is the bias

func LinearForward(x, W, b [][]float64) [][]float64 {
	if !CanMultiply(x, W) {
		fmt.Println("linear forward: cannot multiply x and W (check dimensions)")
		return [][]float64{}
	}

	prod := Multiply(x, W)
	if len(prod) == 0 {
		return prod
	}

	if !CanAddOrSubtract(prod, b) {
		fmt.Println("linear forward: bias shape does not match xW output shape (check dimensions)")
		return [][]float64{}
	}

	return Add(prod, b)
}
