// Testing basic math operations, slicing, 

package main;

import (
  "math"
  "math/rand"
	"testing"
)

type Matrix struct {
  rows int
  cols int
  values []float64
}

func randMatrix(rows, cols int) *Matrix {
  matrix := make([]float64, rows*cols);
  for i := range matrix {
    matrix[i] = rand.NormFloat64();
  }
  return &Matrix{rows, cols, matrix};
}

func constMatrix(rows, cols int, v float64) *Matrix {
	matrix := make([]float64, rows*cols);
	for i := range matrix {
	  matrix[i] = v;
	}
  return &Matrix{rows, cols, matrix};
}

func add(a, b *Matrix) *Matrix {
  c := make([]float64, len(a.values));
  for i := range c {
    c[i] = a.values[i] + b.values[i];
  }
  return &Matrix{a.rows, a.cols, c};
}

func mult(a, b *Matrix) *Matrix {
  c := make([]float64, a.rows * b.cols);

  for i := 0; i < a.rows; i++ {
    for j := 0; j < b.cols; j++ {
      sum := 0.0
      for k := 0; k < b.rows; k++ {
        sum += a.values[a.cols * i + k] * b.values[b.cols * k + j];
      }
      c[b.cols * i + j] = sum;
    }
  }
  return &Matrix{a.rows, b.cols, c};
}

func relu(a *Matrix) *Matrix {
  c := make([]float64, a.rows * a.cols);
  for i := range c {
    c[i] = math.Max(0, a.values[i]);
  }
  return &Matrix{a.rows, a.cols, c};
}

// Benchmarking functions
func BenchmarkMatMult(bench *testing.B) {
  a := randMatrix(1000, 1000)
  b := randMatrix(1000, 1000)
  bench.ResetTimer()
  for i := 0; i < bench.N; i++ {
    mult(a, b);
  }
}

func BenchmarkRelu(bench *testing.B) {
  a := randMatrix(1000, 1000)
  bench.ResetTimer()
  for i := 0; i < bench.N; i++ {
    relu(a);
  }
}
