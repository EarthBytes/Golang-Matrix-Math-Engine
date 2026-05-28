package main

import (
	"flag"
	"fmt"
)

func main() {
	op := flag.String("op", "multiply", "operation: add, subtract, multiply, transpose") // command-line flag for operation
	flag.Parse()

	A := [][]float64{
		{1, 2},
		{3, 4},
	}

	B := [][]float64{
		{5, 6},
		{7, 8},
	}

	var result [][]float64

	switch *op {
	case "add":
		result = Add(A, B)
	case "subtract", "sub":
		result = Subtract(A, B)
	case "multiply", "mul":
		result = Multiply(A, B)
	case "transpose", "trans":
		result = Transpose(A)
	default:
		fmt.Printf("unknown operation %q (use add, subtract, multiply, or transpose)\n", *op)
		return
	}

	printMatrix(result)
}

func printMatrix(m [][]float64) {
	if len(m) == 0 {
		return
	}
	for _, row := range m {
		fmt.Println(row)
	}
}
