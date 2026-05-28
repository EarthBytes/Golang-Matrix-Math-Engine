package main

import "errors"

// matrix operations return (result, error) instead of printing inside matrix_ops
var (
	ErrCannotAdd       = errors.New("cannot add the given matrices")
	ErrCannotSubtract  = errors.New("cannot subtract the given matrices")
	ErrCannotMultiply  = errors.New("cannot multiply the given matrices")
	ErrCannotTranspose = errors.New("cannot transpose the given matrix")

	ErrLinearCannotMultiply = errors.New("linear forward: cannot multiply x and W (check dimensions)")
	ErrLinearBiasShape      = errors.New("linear forward: bias shape does not match xW output shape (check dimensions)")
)

// before, operations called fmt.Println on failure and returned an empty matrix
// but empty matrix looks like a valid result

// now, on failure return nil, err prints the error for the user
// on success return the matrix and err == nil
