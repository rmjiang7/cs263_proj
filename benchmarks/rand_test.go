// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main;

import (
	"errors"
	"fmt"
	"math"
	"testing"
  . "math/rand"
)

const (
	numTestSamples = 10000
)

type statsResults struct {
	mean        float64
	stddev      float64
	closeEnough float64
	maxError    float64
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func nearEqual(a, b, closeEnough, maxError float64) bool {
	absDiff := math.Abs(a - b)
	if absDiff < closeEnough { // Necessary when one value is zero and one value is close to zero.
		return true
	}
	return absDiff/max(math.Abs(a), math.Abs(b)) < maxError
}

var testSeeds = []int64{1, 1754801282, 1698661970, 1550503961}

// checkSimilarDistribution returns success if the mean and stddev of the
// two statsResults are similar.
func (this *statsResults) checkSimilarDistribution(expected *statsResults) error {
	if !nearEqual(this.mean, expected.mean, expected.closeEnough, expected.maxError) {
		s := fmt.Sprintf("mean %v != %v (allowed error %v, %v)", this.mean, expected.mean, expected.closeEnough, expected.maxError)
		fmt.Println(s)
		return errors.New(s)
	}
	if !nearEqual(this.stddev, expected.stddev, expected.closeEnough, expected.maxError) {
		s := fmt.Sprintf("stddev %v != %v (allowed error %v, %v)", this.stddev, expected.stddev, expected.closeEnough, expected.maxError)
		fmt.Println(s)
		return errors.New(s)
	}
	return nil
}

func getStatsResults(samples []float64) *statsResults {
	res := new(statsResults)
	var sum, squaresum float64
	for _, s := range samples {
		sum += s
		squaresum += s * s
	}
	res.mean = sum / float64(len(samples))
	res.stddev = math.Sqrt(squaresum/float64(len(samples)) - res.mean*res.mean)
	return res
}

func checkSampleDistribution(t *testing.T, samples []float64, expected *statsResults) {
	t.Helper()
	actual := getStatsResults(samples)
	err := actual.checkSimilarDistribution(expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func checkSampleSliceDistributions(t *testing.T, samples []float64, nslices int, expected *statsResults) {
	t.Helper()
	chunk := len(samples) / nslices
	for i := 0; i < nslices; i++ {
		low := i * chunk
		var high int
		if i == nslices-1 {
			high = len(samples) - 1
		} else {
			high = (i + 1) * chunk
		}
		checkSampleDistribution(t, samples[low:high], expected)
	}
}

//
// Normal distribution tests
//

func generateNormalSamples(nsamples int, mean, stddev float64, seed int64) []float64 {
	r := New(NewSource(seed))
	samples := make([]float64, nsamples)
	for i := range samples {
		samples[i] = r.NormFloat64()*stddev + mean
	}
	return samples
}

// Benchmarks

func BenchmarkInt63Threadsafe(b *testing.B) {
	for n := b.N; n > 0; n-- {
		Int63()
	}
}

func BenchmarkInt63Unthreadsafe(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Int63()
	}
}

func BenchmarkIntn1000(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Intn(1000)
	}
}

func BenchmarkInt63n1000(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Int63n(1000)
	}
}

func BenchmarkInt31n1000(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Int31n(1000)
	}
}

func BenchmarkFloat32(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Float32()
	}
}

func BenchmarkFloat64(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Float64()
	}
}

func BenchmarkPerm3(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Perm(3)
	}
}

func BenchmarkPerm30(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Perm(30)
	}
}

func BenchmarkPerm30ViaShuffle(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		p := make([]int, 30)
		for i := range p {
			p[i] = i
		}
		r.Shuffle(30, func(i, j int) { p[i], p[j] = p[j], p[i] })
	}
}

// BenchmarkShuffleOverhead uses a minimal swap function
// to measure just the shuffling overhead.
func BenchmarkShuffleOverhead(b *testing.B) {
	r := New(NewSource(1))
	for n := b.N; n > 0; n-- {
		r.Shuffle(52, func(i, j int) {
			if i < 0 || i >= 52 || j < 0 || j >= 52 {
				b.Fatalf("bad swap(%d, %d)", i, j)
			}
		})
	}
}

func BenchmarkRead3(b *testing.B) {
	r := New(NewSource(1))
	buf := make([]byte, 3)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func BenchmarkRead64(b *testing.B) {
	r := New(NewSource(1))
	buf := make([]byte, 64)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}

func BenchmarkRead1000(b *testing.B) {
	r := New(NewSource(1))
	buf := make([]byte, 1000)
	b.ResetTimer()
	for n := b.N; n > 0; n-- {
		r.Read(buf)
	}
}
