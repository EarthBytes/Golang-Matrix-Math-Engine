package main

// LinearForward computes y = xW + b
// x is the input, W is the weight matrix, and b is the bias

func LinearForward(x, W, b [][]float64) ([][]float64, error) { // return error if  checks fail
	if !CanMultiply(x, W) {
		return nil, ErrLinearCannotMultiply
	}

	prod, err := Multiply(x, W) // compute xW and check for error
	if err != nil {
		return nil, err
	}

	if !CanAddOrSubtract(prod, b) {
		return nil, ErrLinearBiasShape
	}

	return Add(prod, b)
}
