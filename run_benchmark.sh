#!/bin/bash

COUNT=5
BENCHMEM=true
GCCFLAGS='-O3 -march=native'

# Turn garbage collection settings
# Default is 100, off is 0
GOGC=100

echo "Running Basic Benchmarks"
cd benchmarks
GOGC=$GOGC go test go_basic_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gc/basic_go
sleep 60
GOGC=$GOGC go test go_basic_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gccgo/basic_go
sleep 60

GOGC=$GOGC go test pure_go_matrix_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gc/matrix
sleep 60
GOGC=$GOGC go test pure_go_matrix_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gccgo/matrix
sleep 60

GOGC=$GOGC go test math_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gc/math
sleep 60
GOGC=$GOGC go test math_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gccgo/math
sleep 60

GOGC=$GOGC go test rand_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gc/rand
sleep 60
GOGC=$GOGC go test rand_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gccgo/rand
sleep 60

GOGC=$GOGC go test strings_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gc/strings
sleep 60
GOGC=$GOGC go test strings_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gccgo/strings
sleep 60

GOGC=$GOGC go test encode_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gc/encoding
sleep 60
GOGC=$GOGC go test encode_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gccgo/encoding
sleep 60

GOGC=$GOGC go test bzip2_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gc/bzip2
sleep 60
GOGC=$GOGC go test bzip2_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gccgo/bzip2
sleep 60

GOGC=$GOGC go test search_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gc/search
sleep 60
GOGC=$GOGC go test search_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > /home/rmjiang/cs263_proj/results/gccgo/search
