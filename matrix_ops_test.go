package main

import (
	"errors"
	"testing"
)

// check two matrices have the same numbers
func sameMatrix(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestMultiply2x2(t *testing.T) {
	A := [][]float64{{1, 2}, {3, 4}}
	B := [][]float64{{5, 6}, {7, 8}}

	got, err := Multiply(A, B)
	if err != nil {
		t.Fatal(err)
	}

	// known result from hand calculation / demo
	want := [][]float64{{19, 22}, {43, 50}}
	if !sameMatrix(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMultiplyParallelSameAsSerial(t *testing.T) {
	A := [][]float64{{1, 2}, {3, 4}}
	B := [][]float64{{5, 6}, {7, 8}}

	serial, err := Multiply(A, B)
	if err != nil {
		t.Fatal(err)
	}

	parallel, err := MultiplyParallel(A, B)
	if err != nil {
		t.Fatal(err)
	}

	if !sameMatrix(serial, parallel) {
		t.Fatalf("parallel %v != serial %v", parallel, serial)
	}
}

func TestLinearForward(t *testing.T) {
	// same numbers as runLinearDemo in main.go
	x := [][]float64{{1, 2, 3}}
	W := [][]float64{
		{0.1, 0.2},
		{0.3, 0.4},
		{0.5, 0.6},
	}
	b := [][]float64{{0.01, 0.02}}

	got, err := LinearForward(x, W, b)
	if err != nil {
		t.Fatal(err)
	}

	want := [][]float64{{2.21, 2.82}}
	if !sameMatrix(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMultiplyBadShapes(t *testing.T) {
	A := [][]float64{{1, 2}}
	B := [][]float64{{1}, {2}, {3}} // inner dimensions do not match

	_, err := Multiply(A, B)
	if !errors.Is(err, ErrCannotMultiply) {
		t.Fatalf("expected ErrCannotMultiply, got %v", err)
	}
}

func TestAddBadShapes(t *testing.T) {
	A := [][]float64{{1, 2}}
	B := [][]float64{{1, 2, 3}}

	_, err := Add(A, B)
	if !errors.Is(err, ErrCannotAdd) {
		t.Fatalf("expected ErrCannotAdd, got %v", err)
	}
}

// build a rows x cols matrix filled with simple values (for benchmarks only)
func benchMatrix(rows, cols int) [][]float64 {
	m := make([][]float64, rows)
	v := 0.1
	for i := range m {
		m[i] = make([]float64, cols)
		for j := range m[i] {
			m[i][j] = v
			v += 0.001
		}
	}
	return m
}

func BenchmarkMultiply_Serial(b *testing.B) {
	A := benchMatrix(256, 256)
	B := benchMatrix(256, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Multiply(A, B)
	}
}

func BenchmarkMultiply_Parallel(b *testing.B) {
	A := benchMatrix(256, 256)
	B := benchMatrix(256, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MultiplyParallel(A, B)
	}
}
