package main

// reports whether M is non-empty and rectangular (no jagged rows).
func IsValidMatrix(M [][]float64) bool {
	if len(M) == 0 || len(M[0]) == 0 {
		return false
	}
	cols := len(M[0])
	for _, row := range M {
		if len(row) != cols {
			return false
		}
	}
	return true
}

// reports whether A and B have the same shape for add/subtract.
func CanAddOrSubtract(A, B [][]float64) bool {
	if !IsValidMatrix(A) || !IsValidMatrix(B) {
		return false
	}
	return len(A) == len(B) && len(A[0]) == len(B[0])
}

// reports whether A and B can be multiplied
func CanMultiply(A, B [][]float64) bool {
	if !IsValidMatrix(A) || !IsValidMatrix(B) {
		return false
	}
	return len(A[0]) == len(B)
}
