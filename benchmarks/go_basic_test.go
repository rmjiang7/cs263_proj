// Testing basic operations in go

package main;

import (
  "fmt"
	"testing"
)

func add_N(N int) {
  var s int = 1 
  for i := 0; i < N; i++ {
    s += (i + 1)
  }
}

func benchmarkGoroutines(N int, bench *testing.B) {
  bench.ResetTimer()
  for n := 0; n < bench.N; n++ {
    for i := 0; i < N; i++ {
      go add_N(N - i)
    }
  }
}

func BenchmarkGoroutines5(bench *testing.B) { benchmarkGoroutines(5, bench) }
func BenchmarkGoroutines10(bench *testing.B) { benchmarkGoroutines(10, bench) }
func BenchmarkGoroutines100(bench *testing.B) { benchmarkGoroutines(100, bench) }

var global int
func benchmarkLoops(N int, bench *testing.B) {
  bench.ResetTimer()
  for n:= 0; n < bench.N; n++ {
    for i := 0; i < N; i++ {
      global += i 
    }
  }
  global += 1
}

func BenchmarkLoops100(bench *testing.B) { benchmarkLoops(100, bench) }
func BenchmarkLoops1000(bench *testing.B) { benchmarkLoops(1000, bench) }
func BenchmarkLoops10000(bench *testing.B) { benchmarkLoops(10000, bench) }
func BenchmarkLoops100000(bench *testing.B) { benchmarkLoops(100000, bench) }

func callTrivialFunction() {}
func BenchmarkFunctionCalling(bench *testing.B) {
  bench.ResetTimer()
  for n:= 0; n < bench.N; n++ {
    callTrivialFunction()
  }
}

func EscapeOptimizedFunctionSum() int {
  numbers := make([]int, 100)
  for i := range numbers {
    numbers[i] = i + 1
  }
  var sum int
  for _, i := range numbers {
    sum += i
  }
  return sum
}

func EscapeFunctionNumbers() []int {
  numbers := make([]int, 100)
  for i := range numbers {
    numbers[i] = i + 1
  }
  return numbers
}

func BenchmarkAllocationEscapeNotOptimized(bench *testing.B) {
  bench.ResetTimer()
  for n:= 0; n < bench.N; n++ {
    numbers := EscapeFunctionNumbers()
    var sum int
    for _, i := range numbers {
      sum += i
    }
  }
}

func BenchmarkAllocationFull(bench *testing.B) {
  bench.ResetTimer()
  var global int
  for n:= 0; n < bench.N; n++ {
    for k := 0; k < 100000; k++ {
      numbers := EscapeFunctionNumbers()
      var sum int
      for _, i := range numbers {
        sum += i
      }
      global += sum
    }
  }
  fmt.Println(global)
}

func BenchmarkAllocationEscapeOptimized(bench *testing.B) {
  bench.ResetTimer()
  var global int
  for n:= 0; n < bench.N; n++ {
    for k := 0; k < 100000; k++ {
      sum := EscapeOptimizedFunctionSum()
      global += sum
    }
  }
  fmt.Println(global)
}

