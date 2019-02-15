#!/bin/bash

COUNT=5
BENCHMEM=true
GCCFLAGS='-O3'

# Turn garbage collection settings
# Default is 100, off is 0
GOGC=100

GOGC=$GOGC go test ./benchmarks -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc > results/gc_benchmarks
GOGC=$GOGC go test ./benchmarks -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags=$GCCFLAGS > results/gccgo_benchmarks


