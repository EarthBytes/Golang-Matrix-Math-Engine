package main

import (
	"flag"
	"fmt"
)

func main() {
	demo := flag.String("demo", "linear", "demo mode: linear (y=xW+b) or op (single operation)")                // chooses the mode - defaults to linear
	op := flag.String("op", "multiply", "when -demo=op: add, subtract, multiply, multiply-parallel, transpose") // chooses the specific math operation - defaults to multiply
	flag.Parse()

	switch *demo { // determines which demo to run based on the -demo flag
	case "linear":
		runLinearDemo()
	case "op":
		runOpDemo(*op)
	default:
		fmt.Printf("unknown demo %q (use linear or op)\n", *demo)
	}
}

func runLinearDemo() {
	x := [][]float64{{1, 2, 3}}
	W := [][]float64{
		{0.1, 0.2},
		{0.3, 0.4},
		{0.5, 0.6},
	}
	b := [][]float64{{0.01, 0.02}}

	fmt.Println("Linear layer forward pass: y = xW + b")
	fmt.Println()
	fmt.Printf("x = %v  (1 × in_features)\n", x)
	fmt.Printf("W = %v  (in_features × out_features)\n", W)
	fmt.Printf("b = %v  (1 × out_features)\n", b)
	fmt.Println()

	y, err := LinearForward(x, W, b)
	if err != nil { // check for error in linear forward pass and print if exists
		fmt.Println(err)
		return
	}
	fmt.Printf("y = %v  (1 × out_features)\n", y)
}

func runOpDemo(op string) {
	A := [][]float64{
		{1, 2},
		{3, 4},
	}

	B := [][]float64{
		{5, 6},
		{7, 8},
	}

	var result [][]float64
	var err error

	// runOpDemo function executes based on which matrix operation is requested
	switch op {
	case "add":
		result, err = Add(A, B)
		if err != nil { // check for error in addition and print if exists
			fmt.Println(err)
			return
		}
		fmt.Println("Matrix addition (A + B). Result:")

		// loop through the matrix and print each row on a new line
		for _, row := range result {
			fmt.Println(row)
		}

	case "subtract", "sub":
		result, err = Subtract(A, B)
		if err != nil { // check for error in subtraction and print if exists
			fmt.Println(err)
			return
		}
		fmt.Println("Matrix subtraction (A - B) Result:")
		for _, row := range result {
			fmt.Println(row)
		}

	case "multiply", "mul":
		result, err = Multiply(A, B)
		if err != nil { // check for error in multiplication and print if exists
			fmt.Println(err)
			return
		}
		fmt.Println("Matrix multplication (A * B) Result:")
		for _, row := range result {
			fmt.Println(row)
		}

	case "multiply-parallel", "mul-par":
		result, err = MultiplyParallel(A, B)
		if err != nil { // check for error in parallel multiplication and print if exists
			fmt.Println(err)
			return
		}
		fmt.Println("Matrix multiplication (A * B) with parallel rows. Result:")
		for _, row := range result {
			fmt.Println(row)
		}

	case "transpose", "trans":
		result, err = Transpose(A)
		if err != nil { // check for error in transposition and print if exists
			fmt.Println(err)
			return
		}
		fmt.Println("Matrix transpose (A^T) Result:")
		for _, row := range result {
			fmt.Println(row)
		}

	default:
		fmt.Printf("unknown operation %q\n", op)
		return
	}
}

func printMatrix(m [][]float64) {
	if len(m) == 0 {
		return
	}
	for _, row := range m {
		fmt.Println(row)
	}
}
