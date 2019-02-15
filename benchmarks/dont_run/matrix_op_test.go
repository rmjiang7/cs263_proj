// Testing matrix math operations that are common in Machine Learning
//
// Original code from:
// https://github.com/aligusnet/linalg-benchmarks/blob/master/gonum-matrix/benchmarks_test.go#L116
//
// Modified to work with current versions of gonum and benchmarking tools
// TODO: This currently does not compile in GCCGO... 


package main;

import (
	"math/rand"
	"testing"
	"gonum.org/v1/gonum/mat"
)

func randMatrix(rows, cols int) *mat.Dense {
	data := make([]float64, rows*cols)
	for i := range data {
	  data[i] = rand.NormFloat64()
	}
	return mat.NewDense(rows, cols, data)
}

func constMatrix(rows, cols int, v float64) *mat.Dense {
	data := make([]float64, rows*cols)
	for i := range data {
	  data[i] = v
	}
	return mat.NewDense(rows, cols, data)
}

func sigmoid(z *mat.Dense) *mat.Dense {
	rows, cols := z.Dims()
	one := constMatrix(rows, cols, 1)
	var tmp mat.Dense
	res := &tmp
	res.Scale(-1, z)
	res.Exp(res)
	res.Add(one, res)
	res.DivElem(one, res)
	return res
}

func relu(x *mat.Dense) *mat.Dense {
	reluElem := func(i, j int, v float64) float64 {
	  if v < 0 {
			return 0
	  } else {
			return v
	  }
	}
	var tmp mat.Dense
	res := &tmp
	res.Apply(reluElem, x)
	return res
}

func BenchmarkRelu(bench *testing.B) {
	a := randMatrix(1000, 1000)
	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
	  relu(a)
	}
}

func BenchmarkSigmoid(bench *testing.B) {
	a := randMatrix(1000, 1000)
	bench.ResetTimer()
	for i := 0; i < bench.N; i++ {
	  sigmoid(a)
	}
}
