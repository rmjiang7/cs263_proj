#!/bin/bash

COUNT=5
BENCHMEM=true
GCCFLAGS='-O3'

go test ./... -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc > results/gc_benchmarks
go test ./... -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags=$GCCFLAGS > results/gccgo_benchmarks


