// Testing basic math operations, slicing, 

package main;

import (
  "math/rand"
	"testing"
)

func add_N(N int) {
  s := 1 
  for i := 0; i < N; i++ {
    s += (i + 1)
  }
}

func BenchmarkGoroutines(bench *testing.B) {
  bench.ResetTimer()
  for j := 0; j < bench.N; j++ {
    for i := 0; i < 100; i++ {
      go add_N(rand.Intn(100) + 10);
    }
  }
}

