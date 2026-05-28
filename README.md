# Golang Matrix Math Engine

A matrix computation engine in Go, built to explore the core operation behind neural network fully connected (linear) layers:

```
y = xW + b
```

This mirrors the forward pass used in frameworks like PyTorch (`nn.Linear`), Keras (`Dense`), and TensorFlow (`Dense`).

## What this project does

This project implements a small matrix math pipeline, including:

- Shape validation to prevent invalid operations before computation
- Core matrix operations: addition, subtraction, multiplication, transpose
- Parallel matrix multiplication using goroutines (one goroutine per result row)
- A linear layer (`LinearForward`) implementing `xW + b`
- Structured error handling using Go’s `(matrix, error)` return pattern
- A CLI for running demos and experimenting with operations

## Requirements

- Go 1.26+

## Run

**Linear layer demo (default)**

```bash
go run .
```

**Single matrix operations** (built-in 2×2 examples `A` and `B` in `main.go`)

```bash
go run . -demo op -op multiply
go run . -demo op -op multiply-parallel
go run . -demo op -op add
go run . -demo op -op subtract
go run . -demo op -op transpose
```

Aliases: `mul`, `mul-par`, `sub`, `trans`.

## Example output (linear demo)

```
Linear layer forward pass: y = xW + b

x = [[1 2 3]]  (1 × in_features)
W = [[0.1 0.2] [0.3 0.4] [0.5 0.6]]  (in_features × out_features)
b = [[0.01 0.02]]  (1 × out_features)

y = [[2.21 2.82]]  (1 × out_features)
```

## Testing

```bash
go test ./...
go test -v ./...
```

## Benchmarks (serial vs parallel multiply)

On large matrices (256×256 in tests), parallel row multiply is faster. Tiny 2×2 demos are too small for goroutines to help.

```bash
go test -bench=BenchmarkMultiply -benchmem
```

Example run (Apple M2, 256×256 matrices):

```
BenchmarkMultiply_Serial-8      61    19036178 ns/op    530816 B/op    257 allocs/op
BenchmarkMultiply_Parallel-8   316     3915957 ns/op    565873 B/op    772 allocs/op
```

Parallel is about 5× faster per op here (19 ms vs 4 ms). It uses more allocations (772 vs 257) because of goroutine overhead — worthwhile at this size, not on tiny 2×2 demos.

## Project layout

```
matrix-engine/
  main.go              # CLI and demos
  matrix_validate.go   # shape checks
  matrix_ops.go        # operations + parallel multiply
  matrix_linear.go     # LinearForward
  matrix_errors.go     # shared errors
  matrix_ops_test.go   # tests and benchmarks
  go.mod
```

## Possible next steps

- ReLU (or other activation) after the linear layer
- Batched inputs (`N×in` @ `in×out`)
- Backpropagation
