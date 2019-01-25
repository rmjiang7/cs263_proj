
// Use goroutines to estimate PI and compute MC error

package main

import (
  "fmt"
  "sync"
  "math/rand"
  "math"
)

// Compute MC estimate of PI
func estimate_pi(n_draws int, wg *sync.WaitGroup, estimates chan float64) {
  defer wg.Done()
  
  var count float64 = 0.0
  for i:= 0; i < n_draws; i++ {
    x_draw := rand.Float64()
    y_draw := rand.Float64()
    if math.Pow(x_draw,2) + math.Pow(y_draw, 2) <= 1 {
      count += 1.0
    }
  }
  estimates <- 4 * count/float64(n_draws)
}

func main() {

  // Concurrent estimates of pi
  var wg sync.WaitGroup
  
  var n_estimates int = 1000
  var n_draws int = 1000
  
  wg.Add(n_estimates)
  estimates := make(chan float64, n_estimates)
  
  for i := 0; i < n_estimates; i++ {
    go estimate_pi(n_draws, &wg, estimates)
  }

  wg.Wait()
  close(estimates)
  
  // Print statistics 
  var values []float64
  for est := range estimates {
    values = append(values, est)
  }
  
  var sum, mean, sd float64

  for i := 0; i < n_estimates; i++ {
    sum += values[i]
  }
  
  mean = sum/float64(n_estimates)
  for i := 0; i < n_estimates; i++ {
    sd += math.Pow(values[i] - mean, 2)
  }
  sd = math.Sqrt(sd/float64(n_estimates))
  
  fmt.Printf("N estimate : %d, N samples per estimate : %d\n", n_estimates, n_draws)
  fmt.Printf("Mean: %f, Std: %f\n", mean, sd)
}
